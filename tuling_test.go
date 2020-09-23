package tuling

import(
	"context"
	"testing"
)

func TestTuling(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)
	r, _, err := client.DefaultApi.Chat(context.Background(), ChatReq{
		UserInfo: UserInfo{
			ApiKey: "90a7053157900a6089c0c289b27a9a77",
			UserId: "1",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(r)
}