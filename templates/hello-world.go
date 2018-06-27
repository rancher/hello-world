package templates

const (
	reqInfoID = "reqInfo"

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
    <h1>Hello world!</h1>
    <h3>My hostname is {{.Hostname}}</h3>`

	webServices = `{{- $length := len .Services }} 
  {{- if gt $length 0 }}
    <div id='Services'>
      <h3>k8s services found {{$length}}</h3>
    {{ range $k,$v := .Services }}
      <b>{{ $k }}</b> {{ $v }}<br />
    {{ end }}
    </div>
    <br />
  {{ end }}`

	webDetails = `    <button class='button' onclick='myFunction()'>Show request details</button>
    <div id="` + reqInfoID + `" style='display:none'>
      <h3>Request info</h3>
      <b>Host:</b> {{.Host}} <br />
      <b>Pod:</b> {{.Hostname}} </b><br />
    {{ range $k,$v := .Headers }}
      <b>{{ $k }}:</b> {{ $v }}<br />
    {{ end }}
    </div>
    <br />`

	webLinks = `    <div id='rancherLinks' class="row social">
      <a class="p-a-xs" href="https://rancher.com/docs"><img src="img/favicon.png" alt="Docs" height="25" width="25"></a>
      <a class="p-a-xs" href="https://slack.rancher.io/"><img src="img/icon-slack.svg" alt="slack" height="25" width="25"></a>
      <a class="p-a-xs" href="https://github.com/rancher/rancher"><img src="img/icon-github.svg" alt="github" height="25" width="25"></a>
      <a class="p-a-xs" href="https://twitter.com/Rancher_Labs"><img src="img/icon-twitter.svg" alt="twitter" height="25" width="25"></a>
      <a class="p-a-xs" href="https://www.facebook.com/rancherlabs/"><img src="img/icon-facebook.svg" alt="facebook" height="25" width="25"></a>
      <a class="p-a-xs" href="https://www.linkedin.com/groups/6977008/profile"><img src="img/icon-linkedin.svg" height="25" alt="linkedin" width="25"></a>
    </div>
    <br />`

	webTail = `    <script>
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

	HelloWorldTemplate = webHead + `
` + webServices + `
` + webLinks + `
` + webDetails + `
` + webTail
)
