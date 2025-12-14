<script>
  import '../app.css';
  import Header from '@/components/Header.svelte';
</script>

<Header></Header>

<div class="container mx-auto">
  <div class="flex flex-col gap-5 justify-center items-center">
    <p class="text-5xl font-bold p-5">全社プロジェクト管理プラットフォーム</p>
    <p class="text-xl p-5">V字モデル対応で、企画からテストまで一元管理</p>
    <button class="btn btn-primary btn-xl">今すぐ始める</button>
  </div>
  <div class="flex gap-5 py-6">
    <div class="card w-full border border-gray-300 shadow">
      <div class="card-body">
        <div class="card-title">チーム横断管理</div>
        <ul class="list">
          <li class="p-3">複数チーム・サービスの進捗を一元管理</li>
          <li class="list-row">チーム別ダッシュボード</li>
          <li class="list-row">サービス横断ビュー</li>
          <li class="list-row">全社レベル進捗管理</li>
        </ul>
      </div>
    </div>
    <div class="card w-full border border-gray-300 shadow">
      <div class="card-body">
        <div class="card-title">V字モデル対応</div>
        <ul class="list">
          <li class="p-3">要件からテストまでのトレーサビリティ</li>
          <li class="list-row">工程間の関連性を明確化</li>
          <li class="list-row">上流・下流の連携</li>
          <li class="list-row">品質管理の徹底</li>
        </ul>
      </div>
    </div>
    <div class="card w-full border border-gray-300 shadow">
      <div class="card-body">
        <div class="card-title">進捗分析</div>
        <ul class="list">
          <li class="p-3">リアルタイムでの進捗把握と分析</li>
          <li class="list-row">ガントチャート</li>
          <li class="list-row">バーンダウトチャート</li>
          <li class="list-row">品質メトリクス</li>
        </ul>
      </div>
    </div>
  </div>
</div>
