<script lang="ts">
  import '@/app.css';

  let activeTab = 'login';

  export let form;

  function switchTab(tab: string) {
    activeTab = tab;
  }
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
      <!-- タブヘッダー -->
      <div class="tabs tabs-boxed mb-4 mx-6">
        <button
          class="tab {activeTab === 'login' ? 'tab-active' : ''}"
          on:click={() => switchTab('login')}
        >
          ログイン
        </button>
        <button
          class="tab {activeTab === 'signup' ? 'tab-active' : ''}"
          on:click={() => switchTab('signup')}
        >
          新規登録
        </button>
      </div>

      <div class="card-body">
        {#if activeTab === 'login'}
          <!-- ログインフォーム -->
          <form method="POST" action="?/login" class="flex flex-col gap-6">
            {#if form?.errors?.general}
              <div class="text-red-500 text-sm mb-2">{form.errors.general}</div>
            {/if}
            <div class="card-title">ログイン</div>
            {#if form?.errors?.email}
              <div class="text-red-500 text-sm mt-1">{form.errors.email[0]}</div>
            {/if}
            <input
              class="input validator w-full"
              type="email"
              name="email"
              required
              placeholder="mail@site.com"
            />
            {#if form?.errors?.password}
              <div class="text-red-500 text-sm mt-1">{form.errors.password[0]}</div>
            {/if}
            <input
              class="input validator w-full"
              type="password"
              name="password"
              required
              placeholder="パスワード"
            />
            <button type="submit" class="btn btn-secondary">ログイン</button>
          </form>
        {:else}
          <!-- 新規登録フォーム -->
          <form method="POST" action="?/signup" class="flex flex-col gap-6">
            {#if form?.errors?.general}
              <div class="text-red-500 text-sm mb-2">{form.errors.general}</div>
            {/if}
            <div class="card-title">新規登録</div>
            {#if form?.errors?.email}
              <div class="text-red-500 text-sm mt-1">{form.errors.email[0]}</div>
            {/if}
            <input
              class="input validator w-full"
              type="email"
              name="email"
              required
              placeholder="mail@site.com"
            />
            {#if form?.errors?.password}
              <div class="text-red-500 text-sm mt-1">{form.errors.password[0]}</div>
            {/if}
            <input
              class="input validator w-full"
              type="password"
              name="password"
              required
              placeholder="パスワード"
            />
            <button type="submit" class="btn btn-secondary">登録</button>
          </form>
        {/if}
      </div>

      <!-- フッターリンク -->
      <span class="text-center">
        {#if activeTab === 'login'}
          アカウントをお持ちでない方は
          <button class="link link-hover" on:click={() => switchTab('signup')}>新規登録</button>
        {:else}
          アカウントをお持ちの方は
          <button class="link link-hover" on:click={() => switchTab('login')}>ログイン</button>
        {/if}
      </span>
    </div>
  </div>
</div>
