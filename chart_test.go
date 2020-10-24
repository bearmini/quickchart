package quickchart

import (
	"encoding/base64"
	"fmt"
	"log"
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
						LineChartDataset{
							Label: String("Foo"),
							Data:  []Data{1, 2},
						},
					},
				},
			},
			Expected: `{"type":"bar","data":{"labels":["Hello world","Test"],"datasets":[{"data":[1,2],"label":"Foo"}]}}`,
		},
		{
			Chart: Chart{
				encoding: String("base64"),
				Type:     "bar",
				Data: ChartData{
					Labels: []Label{"Hello world", "Test"},
					Datasets: []Dataset{
						LineChartDataset{
							Label: String("Foo"),
							Data:  []Data{1, 2},
						},
					},
				},
			},
			Expected: base64.StdEncoding.EncodeToString([]byte(`{"type":"bar","data":{"labels":["Hello world","Test"],"datasets":[{"data":[1,2],"label":"Foo"}]}}`)),
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

func TestGetURL(t *testing.T) {
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
						LineChartDataset{
							Label: String("Foo"),
							Data:  []Data{1, 2},
						},
					},
				},
			},
			Expected: "https://quickchart.io/chart?c=%7B%22type%22%3A%22bar%22%2C%22data%22%3A%7B%22labels%22%3A%5B%22Hello+world%22%2C%22Test%22%5D%2C%22datasets%22%3A%5B%7B%22data%22%3A%5B1%2C2%5D%2C%22label%22%3A%22Foo%22%7D%5D%7D%7D",
		},
		{
			Chart: Chart{
				encoding: String("base64"),
				Type:     "bar",
				Data: ChartData{
					Labels: []Label{"Hello world", "Test"},
					Datasets: []Dataset{
						LineChartDataset{
							Label: String("Foo"),
							Data:  []Data{1, 2},
						},
					},
				},
			},
			Expected: "https://quickchart.io/chart?c=eyJ0eXBlIjoiYmFyIiwiZGF0YSI6eyJsYWJlbHMiOlsiSGVsbG8gd29ybGQiLCJUZXN0Il0sImRhdGFzZXRzIjpbeyJkYXRhIjpbMSwyXSwibGFiZWwiOiJGb28ifV19fQ%3D%3D&encoding=base64",
		},
	}

	for i, data := range testData {
		data := data // capture
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			actual, err := data.Chart.GetURL()
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
