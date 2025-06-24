<script>
    export let urls = []; // 接收的 URL 列表，例如：["https://baidu.com", "https://bilibili.com"]

    // 计算列数（可根据需要调整布局逻辑）
    $: columns = urls.length > 4 ? 3 : urls.length;
    $: gridStyle = `
    grid-template-columns: repeat(${columns}, 1fr);
    gap: 1rem;
  `;
</script>

<div class="browser-container">
    {#if urls.length === 0}
        <div class="empty-tip">暂无需要展示的网页</div>
    {:else}
        <div class="browser-grid" style={gridStyle}>
            {#each urls as url, i}
                <div class="browser-frame">
                    <div class="header">
                        <div class="controls">
                            <span class="close"></span>
                            <span class="minimize"></span>
                            <span class="expand"></span>
                        </div>
                        <div class="url-bar">{url}</div>
                    </div>
                    <iframe
                        src={url}
                        title={`Browser ${i + 1}`}
                        loading="lazy"
                        class="content"
                    ></iframe>
                </div>
            {/each}
        </div>
    {/if}
</div>

<style>
    .browser-container {
        padding: 1rem;
        height: 100vh;
    }

    .empty-tip {
        text-align: center;
        color: #666;
        padding: 2rem;
    }

    .browser-grid {
        display: grid;
        height: 100%;
    }

    .browser-frame {
        background: #fff;
        border-radius: 8px;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
        overflow: hidden;
        display: flex;
        flex-direction: column;
    }

    .header {
        background: #f5f5f5;
        padding: 8px;
        display: flex;
        align-items: center;
        gap: 6px;
        border-bottom: 1px solid #ddd;
    }

    .controls {
        display: flex;
        gap: 6px;
        padding: 0 8px;
    }

    .controls span {
        width: 12px;
        height: 12px;
        border-radius: 50%;
    }

    .close {
        background: #ff5f57;
    }
    .minimize {
        background: #febc2e;
    }
    .expand {
        background: #28c940;
    }

    .url-bar {
        flex: 1;
        background: #fff;
        border-radius: 4px;
        padding: 4px 8px;
        font-size: 0.8em;
        color: #666;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }

    .content {
        flex: 1;
        border: none;
        background: white;
        min-height: 300px;
    }

    /* 响应式布局 */
    @media (max-width: 768px) {
        .browser-grid {
            grid-template-columns: 1fr !important;
        }
    }
</style>
