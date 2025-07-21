<div class="navbar bg-base-100 shadow">
  <div class="navbar-start items-center gap-5">
    <a href="/" class="btn btn-ghost normal-case">
      <img src="/FlowGrid.svg" alt="" width="32">
      <p class="text-black text-xl">FlowGrid</p>
    </a>
  </div>
  <div class="navbar-end">
    {#if user}
      <div class="avatar avatar-placeholder">
        <div class="bg-neutral text-neutral-content w-12 rounded-full">
          <span class="text-lg">{user.name.charAt(0)}</span>
        </div>
      </div>
    {:else}
      <button class="btn btn-primary">ログイン</button>
    {/if}
  </div>
</div>

<script>
const user = {
	name: "shintaro.katsuta",
};
</script>