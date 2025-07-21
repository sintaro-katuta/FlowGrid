<div class="card w-full border border-gray-300 shadow">
  <div class="card-body">
    <h2 class="card-title">{title}</h2>
    <p class="text-5xl font-bold text-center">{count}</p>
    <div class="card-actions justify-end">
      <div class="badge badge-info">
        <p>{diff}</p>
      </div>
    </div>
  </div>
</div>

<script>
export let title;
export let count;
export let diff;
</script>