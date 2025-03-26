package apis

import (
	"encoding"
	"reflect"

	"airplane/internal/errs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"time"
)

func RequestParser[DTO any]() *requestParser[DTO] {
	return &requestParser[DTO]{}
}

type requestParser[DTO any] struct {
	dto DTO
	//bo  BO

	binders []func(ctx *gin.Context) error
}

// Uri 解析 Uri 資料，tag 需要是 uri
func (r *requestParser[DTO]) Uri() *requestParser[DTO] {
	var f = func(ctx *gin.Context) error {
		if err := ctx.ShouldBindUri(&r.dto); err != nil {
			return errs.TraceWrap(err)
		}
		return nil
	}

	r.binders = append(r.binders, f)
	return r
}

// Query 解析 Query 資料，tag 需要是 form
func (r *requestParser[DTO]) Query() *requestParser[DTO] {
	var f = func(ctx *gin.Context) error {
		if err := ctx.ShouldBindQuery(&r.dto); err != nil {
			return errs.TraceWrap(err)
		}

		if err := customBindQueryWithEnum(ctx, &r.dto); err != nil {
			return errs.TraceWrap(err)
		}

		return nil
	}

	r.binders = append(r.binders, f)
	return r
}

func (r *requestParser[DTO]) File() *requestParser[DTO] {
	var f = func(ctx *gin.Context) error {
		if err := ctx.ShouldBindWith(&r.dto, binding.FormMultipart); err != nil {
			return errs.TraceWrap(err)
		}
		return nil
	}

	r.binders = append(r.binders, f)
	return r
}

// Json 解析 JSON 資料，tag 需要是 json
func (r *requestParser[DTO]) Json() *requestParser[DTO] {
	var f = func(ctx *gin.Context) error {
		if err := ctx.ShouldBindJSON(&r.dto); err != nil {
			return errs.TraceWrap(err)
		}
		return nil
	}

	r.binders = append(r.binders, f)
	return r
}

// Bind 執行綁定資料到 DTO
func (r *requestParser[DTO]) Bind(ctx *gin.Context) (dto *DTO, err error) {
	defer func() {
		if e := recover(); e != nil {
			dto = nil
			err = errors.Errorf("%s", e)
		}
	}()

	for _, binder := range r.binders {
		if err := binder(ctx); err != nil {
			return nil, errs.ErrInvalidPayload.Trace(err)
		}
	}

	return &r.dto, nil
}

// customBindQuery
// 自定義解析
// 為了解決 ShouldBindQuery 無法解析 enumer 的問題
func customBindQueryWithEnum(ctx *gin.Context, obj interface{}) error {
	values := ctx.Request.URL.Query()
	return mapping(obj, values, "enum", textUnmarshalerBinder)
}

type binder func(targetType reflect.Type, data string) (reflect.Value, bool)

// 遞歸處理嵌套結構體
func mapping(obj interface{}, values map[string][]string, tag string, binder binder) error {
	val := reflect.ValueOf(obj).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		tagValue := fieldType.Tag.Get(tag)

		if tagValue == "" && fieldType.Anonymous {
			if field.Kind() == reflect.Ptr {
				// 這裡檢查是否是 `time.Time` 型別，並且希望保持它為 nil
				if fieldType.Type.Elem() == reflect.TypeOf(time.Time{}) {
					// 這裡保持 time.Time 指針為 nil，不初始化
					continue
				}

				if field.CanSet() && field.IsNil() {
					// 初始化指針指向的值
					field.Set(reflect.New(fieldType.Type.Elem()))
				}
				field = field.Elem()
			}
			// 如果是物件就在往裡面丟一次
			if field.Kind() == reflect.Struct {
				if field.CanAddr() && field.Addr().CanInterface() {
					if err := mapping(field.Addr().Interface(), values, tag, binder); err != nil {
						return err
					}
				}
				continue
			}
		}

		if tagValue == "-" {
			continue
		}

		queryValues, ok := values[tagValue]
		if !ok {
			continue
		}

		if field.Kind() == reflect.Slice {
			elemType := field.Type().Elem()
			for _, qv := range queryValues {
				if elem, ok := binder(elemType, qv); ok {
					field.Set(reflect.Append(field, elem))
					continue
				}
			}
		} else {
			if elem, ok := binder(field.Type(), queryValues[0]); ok {
				field.Set(elem)
			}
		}
	}

	return nil
}

func textUnmarshalerBinder(targetType reflect.Type, data string) (reflect.Value, bool) {
	if targetType.Kind() == reflect.Ptr {
		targetValue := reflect.New(targetType.Elem()).Elem()
		if result, ok := textUnmarshalerBinder(targetType.Elem(), data); ok {
			targetValue.Set(result)
			return targetValue.Addr(), true
		}

		return targetValue, false
	}

	// 處理有實現 TextUnmarshaler
	targetValue := reflect.New(targetType).Elem()
	if unmarshaler, ok := targetValue.Addr().Interface().(encoding.TextUnmarshaler); ok {
		unmarshaler.UnmarshalText([]byte(data))
		return targetValue, true
	}

	return reflect.Zero(targetType), false
}
