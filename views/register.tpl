<h3>Register</h3>
<form action="/register" method="post">
    {{ .xsrfdata }}
    <div class="form-group">
        <label for="reg_name">Name</label>
        <input name="name" type="text" class="form-control" id="reg_name" aria-ID="nameHelp" value="{{.Users.Name}}">
        {{ if .Errors.Name }}
        <small id="nameHelp" class="form-text text-muted bg-warning">{{ index .Errors.Name 0 }}</small>
        {{ else }}
        <small id="nameHelp" class="form-text text-muted">Any name you'd like to be identified as (may or may not be your real name).</small>
        {{ end }}
    </div>
    <div class="form-group">
        <label for="reg_email">Email ID</label>
        <input name="email" type="email" class="form-control" id="reg_email" aria-ID="emailHelp" value="{{.Users.Email}}">
        {{ if .Errors.Email }}
        <small id="emailHelp" class="form-text text-muted bg-warning">{{ index .Errors.Email 0 }}</small>
        {{ else }}
        <small id="emailHelp" class="form-text text-muted">Use a real one. Will be used for verification/login.</small>
        {{ end }}
    </div>
    <div class="form-group">
        <label for="reg_pwd">Password</label>
        <input name="password" type="password" class="form-control" id="reg_pwd" aria-ID="pwdHelp" value="{{.Users.Password}}">
        {{ if .Errors.Password }}
        <small id="pwdHelp" class="form-text text-muted bg-warning">{{ index .Errors.Password 0 }}</small>
        {{ else }}
        <small id="pwdHelp" class="form-text text-muted">Any string of size 5 to 50. Make it harder to guess.</small>
        {{ end }}
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
</form>