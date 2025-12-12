package models

type Deployment struct {
	Id             string `json:"id"`
	Message        string `json:"message"`
	DeploymentTime string `json:"deploymentTime"`
}

var Deployments = []Deployment{
	{
		Id:             "1",
		Message:        "Deployment 1",
		DeploymentTime: "1:45",
	},
	{
		Id:             "2",
		Message:        "Deployment 2",
		DeploymentTime: "4:15",
	},
	{
		Id:             "3",
		Message:        "Deployment 3",
		DeploymentTime: "7:16",
	},
}
