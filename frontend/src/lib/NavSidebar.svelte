<script>
  import c from "../config"
  // 定义props
  export let items = [];
  export let activeItem = '';
  export let checkActivate = (item)=>{activeItem === item.id};
  
  // 定义事件回调函数
  export let onItemSelect = (item) => {}; // 默认空函数
  
  function handleClick(item, event) {
    event.preventDefault();
    onItemSelect(item); // 调用父组件传递的回调
  }
</script>

<aside class="nav-sidebar">
  <div class="logo">
    <i class="fas fa-code"></i>
    <span>{c.siteTitle}</span>
  </div>
  <nav>
    <ul>
      {#each items as item}
        <li class:active={checkActivate(item)}>
          <a 
            href={item.href} 
            on:click|preventDefault={(e) => handleClick(item, e)}
          >
            {#if item.icon}
              <span class="icon">{item.icon}</span>
            {/if}
            <span class="text">{item.text}</span>
          </a>
        </li>
      {/each}
    </ul>
  </nav>
</aside>

<style>
  .logo {
    top: 0;
    margin-top: 0;
    background-color: #0366d6;
    height: 60px;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
    font-size: 1.5rem;
    font-weight: 700;
  }

  .nav-sidebar {
    background-color: var(--sidebar-bg, #ffffff);
    border-right: 1px solid var(--border-color, #e1e4e8);
    padding: 0 0;
    overflow-y: auto;
    flex-shrink: 0;
    position: sticky;
    top: 0;
  }

  .nav-sidebar nav ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  .nav-sidebar nav li {
    margin: 4px 0;
  }

  .nav-sidebar nav li a {
    display: flex;
    align-items: center;
    padding: 8px 16px;
    color: var(--text-color, #24292e);
    text-decoration: none;
    border-radius: 6px;
    margin: 0 8px;
    transition: all 0.2s ease;
  }

  .nav-sidebar nav li a:hover {
    background-color: var(--hover-bg, #f0f3f6);
  }

  .nav-sidebar nav li.active a {
    background-color: var(--active-bg, #e1f0ff);
    color: var(--active-text, #0366d6);
    font-weight: 500;
  }

  .nav-sidebar .icon {
    margin-right: 8px;
    display: flex;
    align-items: center;
  }

</style>