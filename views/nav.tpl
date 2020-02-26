{{define "navigation"}}
<nav class="navbar navbar-expand-lg navbar-light text-info">
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavAltMarkup" aria-controls="navbarNavAltMarkup" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>
  <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
    <div class="navbar-nav">
        <a class="nav-item nav-link {{ if eq .requestUri "/home"}} active {{ end }}" href="/home">
            Home 
            {{ if eq .requestUri "/home"}} <span class="sr-only">(current)</span> {{ end }}
        </a>
        {{ if .loginId}}
            <a class="nav-item nav-link {{ if eq .requestUri "/user-home"}} active {{ end }}" href="/user-home">
                My Account
                {{ if eq .requestUri "/user-home"}} <span class="sr-only">(current)</span> {{ end }}
            </a>
            <a class="nav-item nav-link {{ if eq .requestUri "/logout"}} active {{ end }}" href="/logout">
                Logout
                {{ if eq .requestUri "/logout"}} <span class="sr-only">(current)</span> {{ end }}
            </a>
        {{ else }}
            <a class="nav-item nav-link {{ if eq .requestUri "/login"}} active {{ end }}" href="/login">
                Login
                {{ if eq .requestUri "/login"}} <span class="sr-only">(current)</span> {{ end }}
            </a>
            <a class="nav-item nav-link {{ if eq .requestUri "/register"}} active {{ end }}" href="/register">
                Register
                {{ if eq .requestUri "/register"}} <span class="sr-only">(current)</span> {{ end }}
            </a>
        {{ end }}
    </div>
  </div>
</nav>
{{end}}


