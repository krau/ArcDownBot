package model

type ArcaeaResponse struct {
	Success bool            `json:"success"`
	Value   arcaeaRespValue `json:"value"`
}

type arcaeaRespValue struct {
	Url     string `json:"url"`
	Version string `json:"version"`
}
