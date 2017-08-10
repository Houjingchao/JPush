package model

type (
	/**
	 /v3/tags/
	 {
      "tags": ["tag1", "tag2"]
     }
	 */
	Tags struct {
		Tags []string `json:"tags"`
	}

	/*
	{
        "registration_ids":{
            "add": [
                "registration_id1",
                "registration_id2"
            ],
            "remove": [
                "registration_id3",
                "registration_id4"
            ]
        }
     }
	 */
	SetTagRequestBody struct {
		SetTages `json:"registration_ids"`
	}
)
