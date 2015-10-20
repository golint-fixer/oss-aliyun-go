package oss

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	accessId   = "3N1H3GnSMozcksyd"
	accessKey  = "u12ckcf2caXVYEEhwjCWOODSbVFx9S"
	testBucket = "testcase"
	testRegion = HangZhou
	testData   = []byte("helloworld")
)

func TestNew(t *testing.T) {
	o := New(testRegion, accessId, accessKey)
	if o == nil {
		t.Error("Unable new oss")
	}
}

func TestPutBucket(t *testing.T) {
	bucket := New(testRegion, accessId, accessKey).Bucket(testBucket)
	err := bucket.PutBucket(PublicRead)
	if err != nil {
		t.Error("Unable put bucket:", err)
	}
}

func TestPut(t *testing.T) {
	bucket := New(testRegion, accessId, accessKey).Bucket(testBucket)
	data := testData
	err := bucket.Put("readme", data, "text/plain", Private)
	if err != nil {
		t.Error("Unable put object:", err)
	}
}

func TestGet(t *testing.T) {
	bucket := New(testRegion, accessId, accessKey).Bucket(testBucket)
	data, err := bucket.Get("readme")
	if err != nil {
		t.Error("Unable get object:", err)
		return
	}
	if bytes.Compare(data, testData) != 0 {
		t.Error("Got wrong object:", err, string(data))
	}
}

func TestURL(t *testing.T) {
	bucket := New(testRegion, accessId, accessKey).Bucket(testBucket)
	url := bucket.URL("readme")
	if url != "http://testcase.oss-cn-hangzhou.aliyuncs.com/readme" {
		t.Error("Unable get correct url:", url)
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		t.Error("Unable get object:", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil || bytes.Compare(data, testData) != 0 {
		t.Error("Got wrong object:", err, string(data))
	}
}

func TestDel(t *testing.T) {
	bucket := New(testRegion, accessId, accessKey).Bucket(testBucket)
	err := bucket.Del("readme")
	if err != nil {
		t.Error("Unable del object:", err)
	}
}

func TestDelBucket(t *testing.T) {
	bucket := New(testRegion, accessId, accessKey).Bucket(testBucket)
	err := bucket.DelBucket()
	if err != nil {
		t.Error("Unable del bucket:", err)
	}
}

func TestPutBuceketWithRegion(t *testing.T) {
	bucket := New(testRegion, accessId, accessKey).Bucket("pinidea-test111")
	err := bucket.PutBucket(PublicRead)
	if err != nil {
		t.Error("Unable put bucket:", err)
	}
}

func TestPutWithRegion(t *testing.T) {
	bucket := New(testRegion, accessId, accessKey).Bucket("pinidea-test111")
	data := []byte("helloworld")
	err := bucket.Put("readme", data, "text/plain", Private)
	if err != nil {
		t.Error("Unable put object:", err)
	}
}

func TestDelWithRegion(t *testing.T) {
	bucket := New(testRegion, accessId, accessKey).Bucket("pinidea-test111")
	if err := bucket.Del("readme"); err != nil {
		t.Error("Unable del with region", err)
	}
}

func TestDelBucketWithRegion(t *testing.T) {
	bucket := New(testRegion, accessId, accessKey).Bucket("pinidea-test111")
	err := bucket.DelBucket()
	if err != nil {
		t.Error("Unable del bucket with region", err)
	}
}
