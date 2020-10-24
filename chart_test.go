package quickchart

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestEncode(t *testing.T) {
	testData := []struct {
		Chart    Chart
		Expected string
	}{
		{
			Chart: Chart{
				Type: "bar",
				Data: ChartData{
					Labels: []Label{"Hello world", "Test"},
					Datasets: []Dataset{
						{
							Label: String("Foo"),
							Data:  []Data{1, 2},
						},
					},
				},
			},
			Expected: url.QueryEscape(`{"type":"bar","data":{"labels":["Hello world","Test"],"datasets":[{"label":"Foo","data":[1,2]}]}}`),
		},
		{
			Chart: Chart{
				encoding: String("base64"),
				Type:     "bar",
				Data: ChartData{
					Labels: []Label{"Hello world", "Test"},
					Datasets: []Dataset{
						{
							Label: String("Foo"),
							Data:  []Data{1, 2},
						},
					},
				},
			},
			Expected: base64.StdEncoding.EncodeToString([]byte(`{"type":"bar","data":{"labels":["Hello world","Test"],"datasets":[{"label":"Foo","data":[1,2]}]}}`)),
		},
	}

	for i, data := range testData {
		data := data // capture
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			actual, err := data.Chart.Encode()
			if err != nil {
				t.Fatalf("unexpected error: %+v\n", err)
			}
			log.Printf("encode: %s\n", actual)
			if actual != data.Expected {
				t.Fatalf("\nExpected: %+v\nActual:   %+v\n", data.Expected, actual)
			}
		})
	}
}
