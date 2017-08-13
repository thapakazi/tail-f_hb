package main

// this one is of projects
func Construct_Url() (url string) {
	url = api_url + "/projects/" + project_id + "/faults?auth_token=" + auth_token + "&q=environment%3A" + env + "&order=" + order + "&limit=" + limit
	// fmt.Println(url)
	return
}
