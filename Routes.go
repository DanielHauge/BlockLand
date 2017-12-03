package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"Join",
		"GET",
		"/join",
		Join,
	},

	Route{
		"Unjoin",
		"DELETE",
		"/join",
		UnJoin,
	},

	Route{
		"Status",
		"GET",
		"/status",
		Status,
	},

	Route{
		"Answer",
		"POST",
		"/answer",
		AnswerOffer,
	},

	Route{
		"Move Out",
		"GET",
		"/moveout",
		MoveOut,
	},
	Route{
		"Chain",
		"GET",
		"/chain",
		GetChain,
	},

}