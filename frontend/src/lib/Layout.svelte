<div class="graceful-layout">
  <slot name="nav"></slot>
  
  <main class="content-area">
    <slot name="context"></slot>
  </main>
</div>

<style>
  .graceful-layout {
    display: grid;
    grid-template-columns: 24vw 76vw;
    grid-template-rows: 100vh;
    background-color: var(--bg-color, #f6f8fa);
  }

  .content-area {
    width: 100%;
    overflow-y: auto;
  }

</style>