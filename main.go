package main

import(
	"context"
	"encoding"
	"fmt"
	"log"
	"net/http"
	"time"
)
type Status string

const(
	UNCHECKED="UNCHECKED"
	UP="UP"
	DOWN="DOWN"
)

var m=map[string]Status{}

type Urls struct{
	Websites []string `json:"websites"`
}

type StatusChecker