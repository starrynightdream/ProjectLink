<!-- MediaViewer.svelte -->
<script>
    import { createEventDispatcher } from 'svelte';
    import MediaLoader from './MediaLoader.svelte';
    import { slide } from 'svelte/transition';

    const dispatch = createEventDispatcher();

    export let paths = [];
    export let currentIndex = 0;
    export let autoPlay = false;
    export let transitionDuration = 500;

    let isTransitioning = false;
    let touchStartX = 0;

    const nextSlide = () => {
        if (isTransitioning || currentIndex >= paths.length - 1) return;
        isTransitioning = true;
        currentIndex++;
        setTimeout(() => isTransitioning = false, transitionDuration);
        dispatch('slide', currentIndex);
    };

    const prevSlide = () => {
        if (isTransitioning || currentIndex <= 0) return;
        isTransitioning = true;
        currentIndex--;
        setTimeout(() => isTransitioning = false, transitionDuration);
        dispatch('slide', currentIndex);
    };

    const handleKeyDown = (e) => {
        if (e.key === 'ArrowRight') nextSlide();
        if (e.key === 'ArrowLeft') prevSlide();
    };
    document.addEventListener("keydown", handleKeyDown);

    const handleTouchStart = (e) => {
        touchStartX = e.touches[0].clientX;
    };

    const handleTouchEnd = (e) => {
        const touchEndX = e.changedTouches[0].clientX;
        const deltaX = touchEndX - touchStartX;
        
        if (deltaX > 50) prevSlide();
        if (deltaX < -50) nextSlide();
    };
</script>

<div
    class="image-browser"
    on:touchstart={handleTouchStart}
    on:touchend={handleTouchEnd}
>
    <!-- 导航按钮 -->
    <div class="navigation">
        <button on:click={prevSlide} class="nav-btn prev" aria-label="上一张">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M15 18L9 12L15 6" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
        </button>
        <!-- <div class="slide-counter">
            {currentIndex + 1}/{paths.length}
        </div> -->
        <button on:click={nextSlide} class="nav-btn next" aria-label="下一张">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M9 18L15 12L9 6" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
        </button>
    </div>

    <!-- 媒体容器 -->
    <div class="media-container">
        {#if paths.length === 0}
            <div class="empty-state">暂无媒体内容</div>
        {:else}
            <div 
                class="media-wrapper"
                style="transform: translateX(-{currentIndex * 100}%)"
            >
                {#each paths as path, index}
                <div class="media-item">
                    <MediaLoader 
                        targetFile={path}
                        on:loaded={() => {}}
                    />
                </div>
                {:else}
                    <div class="loading">加载中...</div>
                {/each}
            </div>
        {/if}
    </div>
    
    <!-- 底部指示器 -->
    <div class="indicators">
        {#each paths as _, index}
            <button
                type="button"
                class="indicator {index === currentIndex ? 'active' : ''}"
                aria-label="选择第 {index + 1} 张"
                aria-current={index === currentIndex}
                on:click={() => currentIndex = index}
            ></button>
        {/each}
    </div>
</div>

<style>
    :global(:root) {
        --primary: #4361ee;
        --primary-light: #4895ef;
        --secondary: #3f37c9;
        --accent: #4cc9f0;
        --light: #f8f9fa;
        --dark: #212529;
        --gray: #6c757d;
        --light-gray: #e9ecef;
        --border: #dee2e6;
        --shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        --shadow-hover: 0 8px 15px rgba(0, 0, 0, 0.2);
        --radius: 8px;
    }

    .image-browser {
        position: relative;
        width: 100%;
        height: 100vh;
        overflow: hidden;
        margin: 0 auto;
        touch-action: pan-y;
        background-color: var(--light);
        border-radius: var(--radius);
        box-shadow: var(--shadow);
    }

    .navigation {
        position: absolute;
        top: 50%;
        width: 100%;
        display: flex;
        justify-content: space-between;
        transform: translateY(-50%);
        padding: 0 20px;
        box-sizing: border-box;
        z-index: 2;
    }

    .nav-btn {
        background: rgba(255, 255, 255, 0.9);
        color: var(--primary);
        border: none;
        padding: 12px;
        border-radius: 50%;
        cursor: pointer;
        font-size: 24px;
        transition: all 0.3s ease;
        display: flex;
        align-items: center;
        justify-content: center;
        width: 48px;
        height: 48px;
        box-shadow: var(--shadow);
        border: 1px solid var(--border);
    }

    .nav-btn:hover {
        background: white;
        color: var(--secondary);
        box-shadow: var(--shadow-hover);
        transform: scale(1.05);
    }

    .nav-btn:active {
        transform: scale(0.95);
    }

    .nav-btn svg {
        width: 24px;
        height: 24px;
    }

    .slide-counter {
        position: absolute;
        left: 50%;
        transform: translateX(-50%);
        color: white;
        background: rgba(33, 37, 41, 0.8);
        padding: 8px 16px;
        border-radius: 20px;
        font-size: 16px;
        font-weight: 600;
        backdrop-filter: blur(4px);
    }

    .media-container {
        position: relative;
        height: 100%;
        display: flex;
        transition: transform 0.5s ease-in-out;
        border-radius: var(--radius);
        overflow: hidden;
    }

    .media-wrapper {
        display: flex;
        width: max-content;
        height: 100%;
    }

    .media-item {
        min-width: 100%;
        height: 100%;
        object-fit: contain;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: var(--light-gray);
    }

    .empty-state {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        color: var(--gray);
        font-size: 24px;
        font-weight: 500;
    }

    .loading {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        color: var(--gray);
        font-size: 18px;
    }

    .indicators {
        position: absolute;
        bottom: 20px;
        left: 50%;
        transform: translateX(-50%);
        display: flex;
        gap: 10px;
        z-index: 2;
    }

    .indicator {
        width: 12px;
        height: 12px;
        border-radius: 50%;
        background-color: rgba(255, 255, 255, 0.5);
        cursor: pointer;
        transition: all 0.3s ease;
        border: 1px solid rgba(0, 0, 0, 0.1);
    }

    .indicator:hover {
        background-color: rgba(255, 255, 255, 0.8);
        transform: scale(1.2);
    }

    .indicator.active {
        background-color: var(--primary);
        transform: scale(1.2);
        box-shadow: 0 0 5px rgba(67, 97, 238, 0.5);
    }

    /* 响应式设计 */
    @media (max-width: 768px) {
        .nav-btn {
            width: 40px;
            height: 40px;
            padding: 8px;
        }
        
        .slide-counter {
            font-size: 14px;
            padding: 6px 12px;
        }
        
        .indicators {
            bottom: 15px;
        }
        
        .indicator {
            width: 10px;
            height: 10px;
        }
    }

    @media (max-width: 480px) {
        .nav-btn {
            width: 36px;
            height: 36px;
        }
        
        .indicators {
            bottom: 10px;
            gap: 8px;
        }
    }
</style>