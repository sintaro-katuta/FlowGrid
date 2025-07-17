<div class="card w-full border border-gray-300 shadow">
  <div class="card-body">
    <div class="badge badge-primary absolute top-5 right-5">{method}</div>
    <div class="card-title">{project}</div>
    <div class="flex justify-between">
      <p class="text-black text-left">進捗率 {progress}%</p>
      <p class="text-gray-400 text-right">期限 2025年8月14日</p>
    </div>
    <progress class="progress w-full h-6 my-3" value={progress} max="100"></progress>
    <div class="card-actions">
      <div class="btn btn-outline">タスク一覧</div>
      <div class="btn btn-outline">ドキュメント</div>
    </div>
  </div>
</div>

<script>
  export let project;
  export let progress
  export let method;
</script>