<script>
//   <MarkdownRenderer target={currentTarget} />
  
//   <!-- 也可以直接指定 -->
//   <!-- <MarkdownRenderer target="README.md" /> -->
  
//   <!-- 如果需要覆盖默认的 API 地址 -->
//   <!-- <MarkdownRenderer target="CHANGELOG.md" apiBaseUrl="http://localhost:10086" /> -->

    import { onMount } from "svelte";
    import "./MarkdownRender.css";
    import { AccessDescript } from "../api/backend"

    export let target = "mock_info.md";

    let htmlContent = "";
    let isLoading = false;
    let error = null;

    async function fetchHtmlContent() {
        isLoading = true;
        error = null;

        try {
            htmlContent = await AccessDescript(target)
        } catch (err) {
            error = err.message;
            console.error("获取内容失败:", err);
        } finally {
            isLoading = false;
        }
    }

    onMount(() => {
        fetchHtmlContent();
    });

    $: if (target) {
        fetchHtmlContent();
    }
</script>

<div class="markdown-container">
    {#if isLoading}
        <div class="loading">
            <div class="spinner"></div>
            加载中...
        </div>
    {:else if error}
        <div class="error">
            <svg class="icon" viewBox="0 0 24 24">
                <path
                    fill="currentColor"
                    d="M12,2C17.53,2 22,6.47 22,12C22,17.53 17.53,22 12,22C6.47,22 2,17.53 2,12C2,6.47 6.47,2 12,2M15.59,7L12,10.59L8.41,7L7,8.41L10.59,12L7,15.59L8.41,17L12,13.41L15.59,17L17,15.59L13.41,12L17,8.41L15.59,7Z"
                />
            </svg>
            {error}
            <button on:click={fetchHtmlContent}>重试</button>
        </div>
    {:else}
        <div class="clear">
        <div class="markdown-body">
            {@html htmlContent}
        </div>
        </div>
    {/if}
</div>

<style>
    .clear {
        all: initial;
    }
    /* 容器样式 */
    .markdown-container {
        max-width: 800px;
        margin: 0 auto;
        padding: 20px;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica,
            Arial, sans-serif;
        line-height: 1.6;
        color: #e1e4e8;
    }

    /* 加载状态 */
    .loading {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        padding: 40px;
        color: #8b949e;
    }

    .spinner {
        border: 4px solid rgba(255, 255, 255, 0.1);
        border-radius: 50%;
        border-top: 4px solid #58a6ff;
        width: 40px;
        height: 40px;
        animation: spin 1s linear infinite;
        margin-bottom: 16px;
    }

    @keyframes spin {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }

    /* 错误状态 */
    .error {
        display: flex;
        align-items: center;
        padding: 16px;
        background-color: rgba(248, 81, 73, 0.1);
        border-radius: 6px;
        color: #f85149;
        margin-bottom: 16px;
    }

    .error .icon {
        width: 20px;
        height: 20px;
        margin-right: 8px;
    }

    .error button {
        margin-left: auto;
        background-color: #238636;
        color: white;
        border: none;
        padding: 6px 12px;
        border-radius: 6px;
        cursor: pointer;
        font-size: 14px;
    }

    .error button:hover {
        background-color: #2ea043;
    }
</style>
