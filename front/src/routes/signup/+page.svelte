<script>
  import '@/app.css';
</script>

<div class="flex justify-center py-12">
  <div class="w-full max-w-md">
    <div class="flex justify-center mb-8">
      <a href="/" class="btn btn-ghost normal-case">
        <img src="/FlowGrid.svg" alt="" width="32" />
        <p class="text-black text-xl">FlowGrid</p>
      </a>
    </div>
    <div class="card border border-gray-300 shadow py-5">
      <div class="card-body">
        <div class="card-title">新規登録</div>
        <input class="input validator w-full" type="email" required placeholder="mail@site.com" />
        <div class="validator-hint">有効なEメールアドレスを入力してください。</div>
        <input class="input validator w-full" type="password" required placeholder="パスワード" />
        <div class="validator-hint">半角英数字で8文字以上入力してください。</div>
        <button class="btn btn-secondary">登録</button>
      </div>
      <span class="text-center">
        アカウントをお持ちの方は
        <a href="/login" class="link link-hover">ログイン</a>
      </span>
    </div>
  </div>
</div>
