document.addEventListener("DOMContentLoaded", () => {
    // Initial load
    loadRotatorServices();
    
    // Polling: Refresh rotator service states every 5 seconds
    setInterval(loadRotatorServices, 5000);
});

// Fetches the rotctld json array along with their live systemd state
async function loadRotatorServices() {
    try {
        const response = await fetch('/api/v1/rotators');
        const rotators = await response.json();
        const tbody = document.querySelector("#rotator-table tbody");
        
        tbody.innerHTML = "";

        rotators.forEach(rot => {
            const isRunning = rot.status === "RUNNING";
            const row = document.createElement("tr");
            row.style.borderBottom = "1px solid #eee";
            row.style.height = "40px";
            
            row.innerHTML = `
                <td><strong>${rot.id}</strong></td>
                <td>${rot.model}</td>
                <td><code>${rot.device}</code></td>
                <td>${rot.port}</td>
                <td>${rot.baudrate} bps</td>
                <td style="color: ${isRunning ? '#2ecc71' : '#e74c3c'}; font-weight: bold;">
                    ${rot.status}
                </td>
                <td>
                    <button 
                        onclick="toggleRotatorService(${rot.id}, '${rot.status}')"
                        style="padding: 5px 10px; cursor: pointer; background-color: ${isRunning ? '#e74c3c' : '#2ecc71'}; color: white; border: none; border-radius: 3px;"
                    >
                        ${isRunning ? 'Stop' : 'Start'}
                    </button>
                </td>
            `;
            tbody.appendChild(row);
        });
    } catch (err) {
        console.error("Failed to fetch rotator services from API:", err);
    }
}

// Triggers the systemd action (start/stop) via the Go proxy execution layer
async function toggleRotatorService(id, currentStatus) {
    const action = currentStatus === "RUNNING" ? "stop" : "start";
    try {
        const response = await fetch(`/api/v1/rotator/${id}/service/${action}`, { method: 'POST' });
        if (response.ok) {
            // Immediate reload after successful command execution
            loadRotatorServices();
        } else {
            console.error(`Backend failed to execution action: ${action} on rotator ${id}`);
        }
    } catch (err) {
        console.error("Network error handling rotator service state shift:", err);
    }
}