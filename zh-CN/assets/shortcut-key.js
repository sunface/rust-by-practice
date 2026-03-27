
// 处理键位和功能绑定关系，执行匹配的快捷键的功能
class KeyBinding {
  constructor(combinationKeys, selectorMap) {
    this.combinationKeys = combinationKeys;
    this.selectorMap = selectorMap;
  }

  matchAction(e) {
    return Object.entries(this.combinationKeys).find(([k, val]) => {
      const { prefixKeys = [], keys } = val;
      const isComposing = !!prefixKeys.find((key) => e[key]);
      return isComposing && keys.includes(e.key);
    });
  }

  handleAction(e) {
    const keyType = this.matchAction(e);
    if (keyType) {
      const actionButton = document.querySelector(this.selectorMap[keyType[0]]);
      if (actionButton) actionButton.click();
    }
  }
}

// 键位绑定器，监听 keydown 事件
class EventBinder {
  constructor(contentId, keyBinding) {
    this.contentDom = document.getElementById(contentId);
    this.keyBinding = keyBinding;
  }

  bindKeydown() {
    if (this.contentDom) {
      this.contentDom.addEventListener('keydown', (e) => {
        const target = e.target;
        if (target && target.classList.contains('ace_text-input')) {
          this.keyBinding.handleAction(e);
        }
      });
    }
  }
}

// 封装初始化逻辑
function initializeKeyBinding() {
  const bindCombinationKeysWithAction = {
    // 重置用 i 键代替，编辑框中内置了很多快捷键，没有合适的键位设置了
    'reset': {
      prefixKeys: ['ctrlKey', 'metaKey'],
      keys: ['i']
    },
    'clip': {
      prefixKeys: ['ctrlKey', 'metaKey'],
      keys: ['c']
    },
    'play': {
      prefixKeys: ['ctrlKey', 'metaKey'],
      keys: ['Enter'],
    }
  };

  const combinationKeysWithSelector = {
    'reset': '.reset-button',
    'clip': '.clip-button',
    'play': '.play-button'
  };

  const keyBinding = new KeyBinding(bindCombinationKeysWithAction, combinationKeysWithSelector);
  const eventBinder = new EventBinder('content', keyBinding);
  eventBinder.bindKeydown();
}

window.addEventListener('load', initializeKeyBinding);
