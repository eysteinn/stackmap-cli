package files

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/eysteinn/stackmap-cli/pkg/apirequest"
	"github.com/eysteinn/stackmap-cli/pkg/global"
	"github.com/spf13/cobra"
	"github.com/tidwall/geojson"
)

// createCmd represents the create command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]
		endpoint, err := url.JoinPath(global.DefaultStackmapUrl, "projects", project, "files")
		if err != nil {
			log.Fatal(err)
		}

		resp := ListResponse{}
		err = apirequest.Get(endpoint, &resp)
		if err != nil {
			log.Fatal(err)
		}

		for idx := range resp.Files {
			/*wmslink, _ := url.JoinPath(global.DefaultStackmapUrl, "services", "wms")

			link := "{{HOST}}/services/wms?map=/mapfiles/{{PROJECT}}/rasters.map&program=mapserv&SERVICE=WMS&VERSION=1.3.0&REQUEST=GetMap&BBOX={{BOUNDINGBOX}}&CRS=EPSG:4326&WIDTH=1024&HEIGHT=768&LAYERS={{PRODUCT}}&STYLES=,&CLASSGROUP=black&FORMAT=image/png&TRANSPARENT=true&TIME={{TIME}}"
			link = strings.ReplaceAll(link, "{{HOST}}", global.DefaultStackmapHost)
			link = strings.ReplaceAll(link, "{{PROJECT}}", project)
			link = strings.ReplaceAll(link, "{{PRODUCT}}", resp.Files[idx].Product)

			timeformat := "2006-01-02T15:04:05"
			timestr := resp.Files[idx].Timestamp.Format(timeformat)
			link = strings.ReplaceAll(link, "{{TIME}}", timestr)
			//link = strings.ReplaceAll(link, "{{BOUNDINGBOX}}", bbox)
			obj, err := geojson.Parse(resp.Files[idx].BoundingBox, geojson.DefaultParseOptions)
			if err != nil {
				log.Fatal(err)
			}
			rect := obj.Rect()
			bbox := fmt.Sprintf("%v,%v,%v,%v", rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
			link = strings.ReplaceAll(link, "{{BOUNDINGBOX}}", bbox)
			fmt.Println(link)
			//60.63023877468341,-38.04788694962596,85.41639837197434,72.08181802540517
			resp.Files[idx].WMS = wmslink*/
			resp.Files[idx].FillWMS(project)
		}
		switch global.Output {
		case "json":
			b, err := json.Marshal(resp)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		case "table":
			log.Fatal("No implemented")
		default:
			for _, file := range resp.Files {
				fmt.Println(file.Filename, "\t", file.Product, "\t", file.UUID, "\t", file.Timestamp)
			}
		}
	},
}

type FileResponse struct {
	UUID        string    `json:"uuid,omitempty"`
	Product     string    `json:"product,omitempty"`
	Filename    string    `json:"filename,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	BoundingBox string    `json:"boundingbox,omitempty"`
	WMS         string    `json:"wms,omitempty"`
}

func (f *FileResponse) FillWMS(project string) {
	link := "{{HOST}}/services/wms?map=/mapfiles/{{PROJECT}}/rasters.map&program=mapserv&SERVICE=WMS&VERSION=1.3.0&REQUEST=GetMap&BBOX={{BOUNDINGBOX}}&CRS=EPSG:4326&WIDTH=1024&HEIGHT=768&LAYERS={{PRODUCT}}&STYLES=,&CLASSGROUP=black&FORMAT=image/png&TRANSPARENT=true&TIME={{TIME}}"
	link = strings.ReplaceAll(link, "{{HOST}}", global.DefaultStackmapHost)
	link = strings.ReplaceAll(link, "{{PROJECT}}", project)
	link = strings.ReplaceAll(link, "{{PRODUCT}}", f.Product)

	timeformat := "2006-01-02T15:04:05"
	timestr := f.Timestamp.Format(timeformat)
	link = strings.ReplaceAll(link, "{{TIME}}", timestr)

	obj, err := geojson.Parse(f.BoundingBox, geojson.DefaultParseOptions)
	if err != nil {
		log.Fatal(err)
	}
	rect := obj.Rect()
	bbox := fmt.Sprintf("%v,%v,%v,%v", rect.Min.Y, rect.Min.X, rect.Max.Y, rect.Max.X)
	link = strings.ReplaceAll(link, "{{BOUNDINGBOX}}", bbox)
	f.WMS = link
}

type ListResponse struct {
	apirequest.ApiResponseSuccess
	Files []FileResponse `json:"files,omitempty"`
}
