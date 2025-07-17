<Header></Header>
<div class="container mx-auto">
  <h1 class="text-3xl font-bold p-5">{team.name}</h1>
  <div class="w-full flex gap-5">
    <DashboardCard title="新規タスク" count="3" diff="前月比+1" />
    <DashboardCard title="進行中タスク" count="15" diff="前月比+1" />
    <DashboardCard title="完了タスク" count="15" diff="前月比+1" />
  </div>
  <div class="card w-full border border-gray-300 shadow my-12">
    <div class="card-body">
      <div class="card-title">進行中プロジェクト</div>
        <div class="flex flex-col gap-5">
          <ProjectCard project="ECサイトフロントエンド" progress="65" method="アジャイル" />
        </div>
      </div>
  </div>
</div>


<script>
import "@/app.css";
import Header from "@/components/Header.svelte";
import DashboardCard from "@/components/DashboardCard.svelte";
import ProjectCard from "@/components/ProjectCard.svelte";


const team = {
  name: "チーム運営"
}
</script>
