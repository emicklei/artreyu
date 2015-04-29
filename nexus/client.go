package nexus

type Client interface{
	RepositoryEndpoint string
}


func New(repoEndpoint string) Client {
	return Client{repoEndpoint}
}

/*
	curl -u $(NEXUS_USER):$(NEXUS_PASSWORD) \
	--upload-file ubanita.tgz \
	https://dev.cloudctrl.com/nexus/content/repositories/snapshots/com/ubanita/sdk/latest/ubanita-1.0.tgz	
*/
func (c Client) UploadFile(srcLocation, destLocation string) error {
	return nil
}

func (c Client) DownloadFile(srcLocation, destLocation string) error {
	return nil
}