<script>
  import Layout from "../lib/Layout.svelte";
  import FlexLoopBtn from "../lib/FlexLoopBtn.svelte";
  import NavSidebar from "../lib/NavSidebar.svelte";
  import MarkdownRender from "../lib/MarkdownRender.svelte";
  import DetailRender from "../lib/DetailRender.svelte";
  import { AccessList } from "../api/backend";
  import { onMount } from "svelte";

  let IsDetail = false;

  // 导航栏相关
  let navItems = [
    {
      id: "0",
      target: "mock_info.md",
      text: "主页",
      icon: "🏠",
      data: {},
    },
    { id: "1", target: "a.md", text: "项目1", icon: "⚙", data: {} },
    { id: "2", target: "b.md", text: "项目2", icon: "⚙", data: {} },
  ];
  let activeNavItem = "0";
  let onItemSelect = (item) => {
    activeNavItem = item.id;
  };

  async function fetchData() {
    try {
      let result = await AccessList();
      if (result) {
        navItems = result.map((item, index) => ({
          id: index.toString(),
          target: item.Description.MDFilePath,
          text: item.Name,
          icon: "⚙",
          data: item,
        }));
      }
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }

  onMount(() => {
    fetchData();
  });
</script>

<Layout>
  <NavSidebar
    items={navItems}
    activeItem={activeNavItem}
    {onItemSelect}
    slot="nav"
  />
  <div slot="context">
    {#if IsDetail}
      <DetailRender
        renderTarget={(() => {
          for (const navItem of navItems) {
            if (navItem.id === activeNavItem) {
              return navItem.data;
            }
          }
          return { Name: "None" };
        })()}
      ></DetailRender>
    {:else}
      <MarkdownRender
        target={(() => {
          for (const navItem of navItems) {
            if (navItem.id === activeNavItem) {
              return navItem.target;
            }
          }
          return "";
        })()}
      />
    {/if}
  </div>
</Layout>
<FlexLoopBtn
  onClick={() => {
    IsDetail = !IsDetail;
  }}
></FlexLoopBtn>
