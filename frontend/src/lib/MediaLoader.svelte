<script>
    import { AccessData } from "../api/backend";
    import { onMount } from "svelte";
    export let targetFile = "never.gif";
    let mediaURL = "";
    let mediaType = "image";
    let isLoading = false;
    let error = null;

    function typeCheck(filename) {
        const typeMap = {
            mp4: 'video/mp4',
            gif: 'image/gif',
            png: 'image/png',
            jpg: 'image/jpeg',
            jpeg: 'image/jpeg',
        }
        return typeMap[filename.split('.').pop()] || 'application/octet-stream';
    }

    async function fetchMedia() {
        isLoading = true;
        error = null;
        if (mediaURL) {
            URL.revokeObjectURL(mediaURL);
        }
        try {
            let type = typeCheck(targetFile);
            const blob = await AccessData(targetFile);
            mediaURL = URL.createObjectURL(
                new Blob([blob], { type })
            );
            mediaType = type == 'video/mp4' ? 'video' : 'image';
        } catch (err) {
            error = err.message;
            console.error("Error fetching media:", err);
        } finally {
            isLoading = false;
        }
    }

    $: if (targetFile) {
        fetchMedia();
    }

    onMount(() => {
        if (targetFile) {
            fetchMedia();
        }
        return () => {
            if (mediaURL) {
                URL.revokeObjectURL(mediaURL);
            }
        };
    });
</script>

<div class="media-container">
  {#if isLoading}
    <div class="loader">加载中...</div>
  
  {:else if error}
    <div class="error">{error}</div>
  
  {:else if mediaURL}
    {#if mediaType === 'video'}
      <video 
        controls 
        src={mediaURL} 
        class="media-element"
      >
      <track kind="captions"/>
        您的浏览器不支持视频标签
      </video>
    
    {:else}
      <img 
        src={mediaURL} 
        alt={targetFile}
        class="media-element"
        loading="lazy"
      />
    {/if}
  {/if}
</div>

<style>
  .media-container {
    max-width: 100%;
    margin: 1rem auto;
    position: relative;
  }

  .loader {
    padding: 1rem;
    text-align: center;
    color: #666;
  }

  .error {
    color: #dc3545;
    padding: 1rem;
    background: #ffeef0;
    border-radius: 4px;
  }

  .media-element {
    max-width: 100%;
    max-height: 80vh;
    display: block;
    margin: 0 auto;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  }

  video {
    background: #000;
  }
</style>