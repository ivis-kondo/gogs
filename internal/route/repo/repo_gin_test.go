package repo

import (
	"reflect"
	"testing"

	"github.com/NII-DG/gogs/internal/context"
	mock_context "github.com/NII-DG/gogs/internal/mocks/context"
	"github.com/golang/mock/gomock"
)

func Test_generateMaDmp(t *testing.T) {
	// TODO: mockの準備

	tests := []struct {
		name                string
		PrepareMockContexts func() context.AbstructContext
		PrepareMockRepoUtil func() AbstructRepoUtil
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateMaDmp(tt.PrepareMockContexts(), tt.PrepareMockRepoUtil())
		})
	}
}

func Test_fetchContentsOnGithub(t *testing.T) {
	wantByte := []byte(`{"name":"maDMP_for_test.ipynb","path":"maDMP_for_test.ipynb","sha":"859552c7e0503b939e70987e097dd2e9d236a99a","size":764,"url":"https://api.github.com/repos/NII-DG/maDMP-template/contents/maDMP_for_test.ipynb?ref=unittest","html_url":"https://github.com/NII-DG/maDMP-template/blob/unittest/maDMP_for_test.ipynb","git_url":"https://api.github.com/repos/NII-DG/maDMP-template/git/blobs/859552c7e0503b939e70987e097dd2e9d236a99a","download_url":"https://raw.githubusercontent.com/NII-DG/maDMP-template/unittest/maDMP_for_test.ipynb","type":"file","content":"ewogImNlbGxzIjogWwogIHsKICAgImNlbGxfdHlwZSI6ICJtYXJrZG93biIs\nCiAgICJtZXRhZGF0YSI6IHt9LAogICAic291cmNlIjogWwogICAgIiMg5Y2Y\n5L2T44OG44K544OI55SobWFETVDjg4bjg7Pjg5fjg6zjg7zjg4hcbiIsCiAg\nICAiXG4iLAogICAgIuOBk+OCjOOBr+WNmOS9k+ODhuOCueODiOOBrueCuuOB\nrm1hRE1Q44OG44Oz44OX44Os44O844OI44Gn44GZ44CC44OG44K544OI57WQ\n5p6c44Gr5b2x6Z+/44KS5Y+K44G844GZ44Gf44KB44CB6Kix5Y+v44Gq44GP\n57eo6ZuG44O75YmK6Zmk44GX44Gq44GE44Gn44GP44Gg44GV44GE44CCIgog\nICBdCiAgfQogXSwKICJtZXRhZGF0YSI6IHsKICAia2VybmVsc3BlYyI6IHsK\nICAgImRpc3BsYXlfbmFtZSI6ICJQeXRob24gMyAoaXB5a2VybmVsKSIsCiAg\nICJsYW5ndWFnZSI6ICJweXRob24iLAogICAibmFtZSI6ICJweXRob24zIgog\nIH0sCiAgImxhbmd1YWdlX2luZm8iOiB7CiAgICJjb2RlbWlycm9yX21vZGUi\nOiB7CiAgICAibmFtZSI6ICJpcHl0aG9uIiwKICAgICJ2ZXJzaW9uIjogMwog\nICB9LAogICAiZmlsZV9leHRlbnNpb24iOiAiLnB5IiwKICAgIm1pbWV0eXBl\nIjogInRleHQveC1weXRob24iLAogICAibmFtZSI6ICJweXRob24iLAogICAi\nbmJjb252ZXJ0X2V4cG9ydGVyIjogInB5dGhvbiIsCiAgICJweWdtZW50c19s\nZXhlciI6ICJpcHl0aG9uMyIsCiAgICJ2ZXJzaW9uIjogIjMuOC4xMiIKICB9\nCiB9LAogIm5iZm9ybWF0IjogNCwKICJuYmZvcm1hdF9taW5vciI6IDIKfQo=\n","encoding":"base64","_links":{"self":"https://api.github.com/repos/NII-DG/maDMP-template/contents/maDMP_for_test.ipynb?ref=unittest","git":"https://api.github.com/repos/NII-DG/maDMP-template/git/blobs/859552c7e0503b939e70987e097dd2e9d236a99a","html":"https://github.com/NII-DG/maDMP-template/blob/unittest/maDMP_for_test.ipynb"}}`)

	// モックの呼び出しを管理するControllerを生成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCtx := mock_context.NewMockAbstructContext(mockCtrl)
	dummyData := make(map[string]interface{})

	type args struct {
		blobPath string
		apiToken string
	}
	tests := []struct {
		name                string
		args                args
		want                []byte
		wantErr             bool
		PrepareMockContexts func() context.AbstructContext
	}{
		{
			name: "succeed fetch blob",
			args: args{
				blobPath: "https://api.github.com/repos/NII-DG/maDMP-template/contents/maDMP_for_test.ipynb?ref=unittest",
				apiToken: "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",
			},
			want:    wantByte,
			wantErr: false,
			PrepareMockContexts: func() context.AbstructContext {
				mockCtx.EXPECT().CallData().AnyTimes().Return(dummyData)
				return mockCtx
			},
		},
		{
			name: "failed fetch blob",
			args: args{
				blobPath: "https://api.github.com/repos/no-exists/maDMP-template/contents/maDMP_for_test.ipynb",
				apiToken: "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",
			},
			want:    nil,
			wantErr: true,
			PrepareMockContexts: func() context.AbstructContext {
				mockCtx.EXPECT().CallData().AnyTimes().Return(dummyData)
				return mockCtx
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f repoUtil
			got, err := f.fetchContentsOnGithub(tt.PrepareMockContexts(), tt.args.blobPath, tt.args.apiToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchContentsOnGithub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				// if !bytes.Equal(got, wantByte) {
				t.Errorf("fetchContentsOnGithub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeBlobContent(t *testing.T) {
	rightBlobInfo := []byte(`{"content":"SGVsbG8sIHdvcmxkLg=="}`)
	rightDecordedBlob := "Hello, world."

	wrongJsonInfo := []byte(`{"content":"SGVsbG8sIHdvcmxkLg=="`)

	type args struct {
		blobInfo []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "SucceedDecording",
			args: args{
				blobInfo: rightBlobInfo,
			},
			want:    rightDecordedBlob,
			wantErr: false,
		},
		{
			name: "FailUnmarshal",
			args: args{
				blobInfo: wrongJsonInfo,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f repoUtil
			got, err := f.decodeBlobContent(tt.args.blobInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeBlobContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decodeBlobContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
