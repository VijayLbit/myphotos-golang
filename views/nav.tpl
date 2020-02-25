{{define "navigation"}}
<ul class="nav">
    <li class="nav-item">
        <a class="nav-link" href="/home">Home</a>
    </li>
    <li class="nav-item">
        {{ if .loginId}}
        <a class="nav-link" href="/logout">Logout</a>
        {{ else }}
        <a class="nav-link" href="/login">Login</a>
        {{ end }}
    </li>
    <li class="nav-item">
        <a class="nav-link" href="/register">Register</a>
    </li>
</ul>
{{end}}