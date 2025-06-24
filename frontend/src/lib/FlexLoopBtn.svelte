<script>
  // 接收点击事件处理函数
  export let onClick = () => {};
  
  let isSpinning = false;
  
  // 处理点击事件
  function handleClick() {
    if (isSpinning) return; // 防止重复点击
    
    isSpinning = true;
    onClick(); // 调用传入的函数
    
    // 1秒后停止旋转
    setTimeout(() => {
      isSpinning = false;
    }, 1000);
  }
</script>

<div class="fab-container">
  <button 
    class="fab-button"
    on:click={handleClick}
    aria-label="Action button"
  >
    <svg 
      class="spinner-icon {isSpinning ? 'spin' : ''}" 
      viewBox="0 0 24 24"
      width="24"
      height="24"
    >
      <!-- 圆形箭头图标 -->
      <path 
        fill="currentColor" 
        d="M12,4V2A10,10 0 0,0 2,12H4A8,8 0 0,1 12,4Z" 
      />
    </svg>
  </button>
</div>

<style>
  /* 固定在右下角 */
  .fab-container {
    position: fixed;
    right: 24px;
    bottom: 24px;
    z-index: 1000;
  }
  
  /* 圆形按钮样式 */
  .fab-button {
    width: 56px;
    height: 56px;
    border-radius: 50%;
    background: #6200ee;
    color: white;
    border: none;
    box-shadow: 0 3px 5px rgba(0,0,0,0.2), 0 3px 6px rgba(0,0,0,0.3);
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    outline: none;
  }
  
  .fab-button:hover {
    background: #7c4dff;
    transform: translateY(-2px);
    box-shadow: 0 4px 8px rgba(0,0,0,0.2), 0 6px 12px rgba(0,0,0,0.3);
  }
  
  .fab-button:active {
    transform: translateY(1px);
    box-shadow: 0 2px 4px rgba(0,0,0,0.2), 0 3px 6px rgba(0,0,0,0.2);
  }
  
  /* 旋转动画 */
  .spinner-icon.spin {
    animation: spin 1s ease-in-out;
  }
  
  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }
</style>