// Copyright 2012-2015 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import "testing"

func TestTermVectorWithId(t *testing.T) {
	client := setupTestClientAndCreateIndex(t)

	tweet1 := tweet{User: "olivere", Message: "Welcome to Golang and Elasticsearch."}

	// Add a document
	indexResult, err := client.Index().
		Index(testIndexName).
		Type("tweet").
		Id("1").
		BodyJson(&tweet1).
		Refresh(true).
		Do()
	if err != nil {
		t.Fatal(err)
	}
	if indexResult == nil {
		t.Errorf("expected result to be != nil; got: %v", indexResult)
	}

	// TermVectors by specifying ID
	field := "Message"
	result, err := client.TermVector(testIndexName, "tweet", "1").
		Fields([]string{field}).
		FieldStatistics(true).
		TermStatistics(true).
		Do()
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("expected to return information and statistics")
	}
	if !result.Found {
		t.Errorf("expected found to be %v; got: %v", true, result.Found)
	}
	if result.Took <= 0 {
		t.Errorf("expected took in millis > 0 got: %v", result.Took)
	}
}

func TestTermVectorWithDoc(t *testing.T) {
	client := setupTestClientAndCreateIndex(t)

	// TermVectors by specifying Doc
	var doc = map[string]string{
		"fullname": "John Doe",
		"text":     "twitter test test test",
	}
	var perFieldAnalyzer = map[string]string{
		"fullname": "keyword",
	}

	result, err := client.TermVector(testIndexName, "tweet", "").
		Doc(doc).
		PerFieldAnalyzer(perFieldAnalyzer).
		FieldStatistics(true).
		TermStatistics(true).
		Do()
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("expected to return information and statistics")
	}
	if !result.Found {
		t.Errorf("expected found to be %v; got: %v", true, result.Found)
	}
	if result.Took <= 0 {
		t.Errorf("expected took in millis > 0 got: %v", result.Took)
	}
}
