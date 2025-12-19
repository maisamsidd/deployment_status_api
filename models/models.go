package models

type Deployment struct {
	Id             string `json:"id"`
	Message        string `json:"message"`
	DeploymentTime string `json:"deploymentTime"`
}

var Deployments = []Deployment{
	{
		Id:             "1",
		Message:        "example 1",
		DeploymentTime: "1:45",
	},
}
