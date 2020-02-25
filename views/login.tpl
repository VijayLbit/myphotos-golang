<h3>Login</h3>
<form action="/login" method="post">
    {{ .xsrfdata }}
    <div class="form-group">
        <label for="cred_email">Email ID</label>
        <input name="email" type="email" class="form-control" id="cred_email" aria-ID="emailHelp">
        {{ if .Errors.Email }}
        <small id="emailHelp" class="form-text text-muted bg-warning">{{ index .Errors.Email 0 }}</small>
        {{ else }}
        <small id="emailHelp" class="form-text text-muted">Email ID used for Registration.</small>
        {{ end }}
    </div>
    <div class="form-group">
        <label for="cred_pwd">Password</label>
        <input name="password" type="password" class="form-control" id="cred_pwd" aria-ID="pwdHelp">
        {{ if .Errors.Password }}
        <small id="pwdHelp" class="form-text text-muted bg-warning">{{ index .Errors.Password 0 }}</small>
        {{ else }}
        <small id="pwdHelp" class="form-text text-muted">Account password - required</small>
        {{ end }}
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
</form>