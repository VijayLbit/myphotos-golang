<h3>Login</h3>
<form action="" method="post">
    {{ .xsrfdata }}
    <div class="form-group">
        <label for="cred_email">Email ID</label>
        <input name="email" type="email" class="form-control" id="cred_email" aria-ID="emailHelp">
        <small id="emailHelp" class="form-text text-muted">Email ID used for Registration.</small>
    </div>
    <div class="form-group">
        <label for="cred_pwd">Password</label>
        <input name="password" type="password" class="form-control" id="cred_pwd">
    </div>
    <button type="submit" class="btn btn-primary">Submit</button>
</form>