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
        <div className="w-[300px] h-[400px] p-4 bg-white">
            <h1 className="text-lg font-bold mb-4">Password Manager</h1>
            <button 
                onClick={onclick}
                className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
            >
                Click me
            </button>
        </div>
    )
}

export default Popup;