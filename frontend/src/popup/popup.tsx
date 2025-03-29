function Popup() {
    const onclick = async () => {
        let [tab] = await chrome.tabs.query({ active: true });
        if (!tab?.id) return;
        
        chrome.scripting.executeScript({
            target: { tabId: tab.id },
            func: () => {
                alert('Hello from my extension');
            },
            args: []
        })
    }

    return (
        <button onClick={onclick}>Click me</button>
    )
}

export default Popup;