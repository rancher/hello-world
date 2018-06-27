package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const (
	webHead = `<html>
	<head>
		<title>Rancher</title>
		<link rel="icon" href="img/favicon.png">
		<style>
			body {
				background-color: white;
				text-align: center;
				padding: 50px;
				font-family: "Open Sans","Helvetica Neue",Helvetica,Arial,sans-serif;
			}
			button {
			    background-color: #0075a8; 
			    border: none;
			    color: white;
			    padding: 15px 32px;
			    text-align: center;
			    text-decoration: none;
			    display: inline-block;
			    font-size: 16px;
			}

			#logo {
				margin-bottom: 40px;
			}
		</style>
	</head>
	<body>
		<img id="logo" src="img/rancher-logo.svg" alt="Rancher logo" width=400 />
		<h1>Hello world!</h1>`

	webTail = `		<div class="row social">
			<a class="p-a-xs" href="https://rancher.com/docs"><img src="img/favicon.png" alt="Docs" height="25" width="25"></a>
			<a class="p-a-xs" href="https://slack.rancher.io/"><img src="img/icon-slack.svg" alt="slack" height="25" width="25"></a>
			<a class="p-a-xs" href="https://github.com/rancher/rancher"><img src="img/icon-github.svg" alt="github" height="25" width="25"></a>
			<a class="p-a-xs" href="https://twitter.com/Rancher_Labs"><img src="img/icon-twitter.svg" alt="twitter" height="25" width="25"></a>
			<a class="p-a-xs" href="https://www.facebook.com/rancherlabs/"><img src="img/icon-facebook.svg" alt="facebook" height="25" width="25"></a>
			<a class="p-a-xs" href="https://www.linkedin.com/groups/6977008/profile"><img src="img/icon-linkedin.svg" height="25" alt="linkedin" width="25"></a>
        </div>	
		<script>
			function myFunction() {
			    var x = document.getElementById("` + reqInfoID + `");
			    if (x.style.display === "none") {
			        x.style.display = "block";
			    } else {
			        x.style.display = "none";
			    }
			}
		</script>
	</body>
</html>`

	reqInfoID = "reqInfo"
)

func getServices() map[string]string {
	k8sServices := make(map[string]string)

	for _, evar := range os.Environ() {
		show := strings.Split(evar, "=")
		regName := regexp.MustCompile("^.*_PORT$")
		regLink := regexp.MustCompile("^(tcp|udp)://.*")
		if regName.MatchString(show[0]) && regLink.MatchString(show[1]) {
			k8sServices[strings.TrimSuffix(show[0], "_PORT")] = show[1]
		}
	}

	return k8sServices
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	k8sServices := getServices()

	fmt.Fprintln(w, webHead)
	fmt.Fprintln(w, "		<h3>My hostname is ", hostname, "</h3>")
	k8sNumber := len(k8sServices)
	if k8sNumber > 0 {
		fmt.Fprintln(w, "		<h3>k8s services found,", k8sNumber, "</h3>")
		for k, v := range k8sServices {
			fmt.Fprintln(w, "		<b>", k, "</b> ", v, "<br />")
		}
	}
	fmt.Fprintln(w, "		<br />")

	fmt.Fprintln(w, "		<button class='button' onclick='myFunction()'>Show details</button>")
	fmt.Fprintln(w, "		<div id='"+reqInfoID+"' style='display:none'>")
	fmt.Fprintln(w, "			<h3>Request info</h3>")
	fmt.Fprintln(w, "			<b>Host:</b> ", r.Host, "<br />")
	fmt.Fprintln(w, "			<b>Pod:</b> ", hostname, "</b><br />")
	for k, v := range r.Header {
		if strings.HasPrefix(k, "X-") {
			fmt.Fprintln(w, "		<b>", k, "</b> ", v, "<br />")
		}
	}
	fmt.Fprintln(w, "		<br /></div>")
	fmt.Fprintln(w, webTail)
}

func main() {
	webPort := os.Getenv("WEB_PORT")
	if webPort == "" {
		webPort = "8080"
	}

	fmt.Println("Running web-test service at", webPort, "port")
	http.HandleFunc("/", handler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(os.Getenv("PWD")))))
	http.ListenAndServe(":"+webPort, nil)
}
