package file

import (
	"fmt"
	"testing"

	"github.com/micro-plat/lib4go/utility"
)

func TestSaveBase64Img(t *testing.T) {
	err := SaveBase64Img(`data:image/jpeg;base64,/9j/4QAYRXhpZgAASUkqAAgAAAAAAAAAAAAAAP/sABFEdWNreQABAAQAAAAKAAD/4QMpaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLwA8P3hwYWNrZXQgYmVnaW49Iu+7vyIgaWQ9Ilc1TTBNcENlaGlIenJlU3pOVGN6a2M5ZCI/PiA8eDp4bXBtZXRhIHhtbG5zOng9ImFkb2JlOm5zOm1ldGEvIiB4OnhtcHRrPSJBZG9iZSBYTVAgQ29yZSA1LjAtYzA2MCA2MS4xMzQ3NzcsIDIwMTAvMDIvMTItMTc6MzI6MDAgICAgICAgICI+IDxyZGY6UkRGIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyI+IDxyZGY6RGVzY3JpcHRpb24gcmRmOmFib3V0PSIiIHhtbG5zOnhtcD0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLyIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bXA6Q3JlYXRvclRvb2w9IkFkb2JlIFBob3Rvc2hvcCBDUzUgV2luZG93cyIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDo2NEVCRDRGMzNCMjMxMUU5OUM0Q0UxRTNCQTBCOEZFMSIgeG1wTU06RG9jdW1lbnRJRD0ieG1wLmRpZDo2NEVCRDRGNDNCMjMxMUU5OUM0Q0UxRTNCQTBCOEZFMSI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOjY0RUJENEYxM0IyMzExRTk5QzRDRTFFM0JBMEI4RkUxIiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOjY0RUJENEYyM0IyMzExRTk5QzRDRTFFM0JBMEI4RkUxIi8+IDwvcmRmOkRlc2NyaXB0aW9uPiA8L3JkZjpSREY+IDwveDp4bXBtZXRhPiA8P3hwYWNrZXQgZW5kPSJyIj8+/+4AIUFkb2JlAGTAAAAAAQMAEAMCAwYAAAqrAAASRQAAG/3/2wCEABQQ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`,
		fmt.Sprintf("/home/yanglei/%s.jpg", utility.GetGUID()))
	if err != nil {
		t.Error(err)
	}
}