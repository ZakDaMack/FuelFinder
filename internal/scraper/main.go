package scraper

import (
	"log/slog"
)

type Job struct {
	Url      string
	Stations []StationDataset
}

func NewScraper(url string, out chan Job) error {
	// read through links and then upload to system
	links, err := GetTableLinks(url)
	if err != nil {
		slog.Error("could not scrape fuel price data", "url", url, "error", err)
		return err
	}

	slog.Info("fetched companies from gov website", "count", len(links))

	// GO CHANNEL PIPELINE
	// stage 1: convert array to job channel
	jobs := createChannel(links)
	// stage 2: collect stations from jobs
	stations := fetchData(jobs)
	// // stage 3: filter data
	filtered := filterData(stations)
	// stage 4: sanitize data
	sanitized := sanitizeData(filtered)

	// send the jobs to the output channel
	for j := range sanitized {
		slog.Debug("pipeline job completed", "url", j.Url, "stations", len(j.Stations))
		out <- j
	}

	return nil
}

func createChannel(items []string) <-chan Job {
	out := make(chan Job)
	go func(arr []string) {
		for _, i := range arr {
			slog.Debug("created new job", "url", i)
			out <- Job{
				Url: i,
			}
		}
		close(out)
	}(items)

	return out
}

func fetchData(in <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range in {
			data, err := ReadJsonFrom(job.Url)
			if err != nil {
				slog.Error("could not read json", "url", job.Url, "error", err)
				continue
			}

			job.Stations = data
			out <- job
		}
		close(out)
	}()

	return out
}

func filterData(in <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range in {
			filtered := make([]StationDataset, 0)
			for _, station := range job.Stations {
				if IsValidItem(&station) {
					filtered = append(filtered, station)
				}
			}

			job.Stations = filtered
			out <- job
		}
		close(out)
	}()

	return out
}

func sanitizeData(in <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range in {
			// Sanitise the data
			for i := range job.Stations {
				cleanStationItem(&job.Stations[i])
			}
			out <- job
		}
		close(out)
	}()
	return out
}

// func uploadData(in <-chan Job, client models.FuelFinderClient) <-chan Job {
// 	out := make(chan Job)
// 	go func() {
// 		for job := range in {
// 			req := &models.StationItems{Items: job.stations}
// 			res, err := client.Upload(context.TODO(), req)
// 			if err != nil {
// 				slog.Error("could not upload station items", "url", job.url, "error", err)
// 				continue
// 			}
// 			job.uploadedStations = res.Count
// 		}
// 		close(out)
// 	}()
// 	return out
// }
