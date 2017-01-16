package models

import "testing"

func TestCaculateChecksum(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"../conf/app.conf", "68134F46DA7854064109B8F532419C7D"},
	}
	for _, c := range cases {
		got := caculateChecksum(c.in)
		if got != c.want {
			t.Errorf("caculateChecksum() : \n want\t %s\n got \t %s", c.want, got)
		}
	}

}

func TestGetReleaseFilesInfo(t *testing.T) {
	folder := "../static/release"
	for _, f := range getLocalReleaseFilesInfo(folder) {
		t.Logf("file path:\t%s\n", f.Filepath)
		t.Logf("file name:\t%s\n", f.Filename)
		t.Logf("checksum:\t%s\n", f.Checksum)
		t.Logf("create time:\t%v\n", f.Created)
	}
}

/*
func TestGetBuildNumberFromFileName(t *testing.T) {
	cases := []struct {
		in   string
		want uint64
	}{
		{"../conf/app.conf", 0},
		{"Host_ReleaseCRC_11301052.hex", 11301052},
		{"Host_Release.hex", 0},
	}
	for _, c := range cases {
		got, _ := getBuildNumberFromFileName(c.in)

			if err != nil {
				t.Errorf("getBuildNumberFromFileName() : ", err.Error())
			}

		if got != c.want {
			t.Errorf("getBuildNumberFromFileName() : \n want\t %d\n got \t %d", c.want, got)
		}
	}
}
*/
func TestCheckFileType(t *testing.T) {
	cases := []struct {
		in   string
		want FileType
	}{
		{"../conf/app.conf", FILETYPE_UNKNOWN},
		{"Host_ReleaseCRC_11301052.hex", FILETYPE_APP},
		{"Host_Release.hex", FILETYPE_APP},
		{"GBox_DSP_flash_0.2.11301530.hex", FILETYPE_DSP},
		{"HostBootCRC_11291722.hex", FILETYPE_BOOT},
	}
	for _, c := range cases {
		var f Filerepo
		f.Filename = c.in
		f.checkFileType()
		got := f.Filetype
		if got != c.want {
			t.Errorf("checkFileType() : \n want\t %d\n got \t %d", c.want, got)
		}
	}

}

func TestFindCRC(t *testing.T) {
	cases := []struct {
		in1  string
		in2  FileType
		want string
	}{
		{"S0030424D4", FILETYPE_DSP, "D42404"},
		{"S3150003FFF0FFFFFFFFC230B800F6000000D8B10000D3", FILETYPE_BOOT, "0000B1D8"},
		{"S3150003FFF0FFFFFFFFC230B800F6000000D8B10000D3", FILETYPE_APP, "0000B1D8"},
	}
	for _, c := range cases {
		got := findCRC(c.in1, c.in2)
		if got != c.want {
			t.Errorf("findCRC() : \n want\t %s\n got \t %s", c.want, got)
		}
	}

}

func TestGetFileInfoFromRemoteRepo(t *testing.T) {
	fl, err := getFileInfoFromRemoteRepo()
	if err != nil {
		t.Logf("getFileInfoFromRemoteRepo():%s", err.Error())
	}
	for _, f := range fl {
		t.Logf("%v", f)
	}
}

func TestSyncReleaseFilesInfo(t *testing.T) {
	t.Errorf("TestSyncReleaseFilesInfo 0 ")
	SyncReleaseFilesInfo()

	t.Errorf("TestSyncReleaseFilesInfo 1 ")

	t.Logf("TestSyncReleaseFilesInfo() ... done\n")
}
