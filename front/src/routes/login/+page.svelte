<div class="container  py-12">
  <div class="navbar-start items-center gap-5">
    <a href="/" class="btn btn-ghost normal-case">
      <img src="/FlowGrid.svg" alt="" width="32">
      <p class="text-black text-xl">FlowGrid</p>
    </a>
  </div>
  <div class="card w-1/2 border border-gray-300 shadow py-5">
    <div class="card-body">
      <div class="card-title">ログイン</div>
      <input class="input validator w-full" type="email" required placeholder="mail@site.com" />
      <div class="validator-hint">有効なEメールアドレスを入力してください。</div>
      <input class="input validator w-full" type="password" required placeholder="パスワード" />
      <div class="validator-hint">半角英数字で8文字以上入力してください。</div>
      <button class="btn btn-secondary">ログイン</button>
    </div>
    <span class="text-center">
      アカウントをお持ちでない方は
      <a href="/signup" class="link link-hover">新規登録</a>
    </span>
  </div>
</div>

<script>
import "@/app.css";
</script>