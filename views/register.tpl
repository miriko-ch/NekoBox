{{template "template/header.tpl" .}}
<form method="post" id="form">
    {{ .xsrfdata }}
    <fieldset class="uk-fieldset">
        <legend class="uk-legend">新用户注册</legend>
        {{if ne .error ""}}
            <div class="uk-alert-danger" uk-alert>
                <a class="uk-alert-close" uk-close></a>
                <p>{{.error}}</p>
            </div>
        {{end}}
        <div class="uk-margin">
            <label class="uk-form-label" for="form-stacked-text">电子邮箱地址</label>
            <input name="email" class="uk-input" type="text" value="{{.email}}">
        </div>
        <div class="uk-margin">
            <label class="uk-form-label" for="form-stacked-text">个性域名 (你的问答箱网址： <code>https://box.lonelyion.com/_/你的域名</code>)</label>
            <input name="domain" class="uk-input" type="text" value="{{.domain}}">
        </div>
        <div class="uk-margin">
            <label class="uk-form-label" for="form-stacked-text">昵称</label>
            <input name="name" class="uk-input" type="text" value="{{.name}}">
        </div>
        <div class="uk-margin">
            <label class="uk-form-label" for="form-stacked-text">密码</label>
            <input type="password" name="password" class="uk-input" type="text">
        </div>
        <div class="uk-margin">
            <label class="uk-form-label" for="form-stacked-text">确认密码</label>
            <input type="password" name="repeat_password" class="uk-input" type="text">
        </div>
        <div class="uk-margin">
            <button type="submit" class="uk-button uk-button-primary g-recaptcha" data-sitekey="{{.recaptcha}}"
                    data-callback="onSubmit">注册
            </button>
        </div>
    </fieldset>
</form>
{{template "template/footer.tpl" .}}