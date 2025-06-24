<script>
    import MultyWindows from "./MultyWindows.svelte";
    import MediaViwer from "./MediaViwer.svelte";
    import MediaLoader from "./MediaLoader.svelte";
    export let renderTarget = {};
</script>

<div>
    <h1 style="color: black;">
        {renderTarget.text}
    </h1>
    {#if `${renderTarget.Type}`.toLowerCase()==="url"}
        <MultyWindows urls={renderTarget.FrontParams.URLs||[]}></MultyWindows>
    {:else if `${renderTarget.Type}`.toLowerCase()==="vedio"}
    <div style="height: 100vh; width:100%;">
        <MediaViwer
            paths={renderTarget.FrontParams.Paths}
            on:slide={(e) => console.log("切换到:", e.detail)}
        ></MediaViwer>
    </div>
        <!-- {#each renderTarget.FrontParams.Paths as p,i}
        <div id={`${i}`}>
            <MediaLoader
                targetFile={p}
            ></MediaLoader>
        </div>
        {/each} -->
    {/if}
</div>