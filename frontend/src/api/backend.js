
const apiBaseUrl = import.meta.env.VITE_API_BASE_URL || "http://localhost:3000";

export const AccessList = async () => {
    const response = await fetch(`${apiBaseUrl}/projects/list`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
    });

    if (!response.ok) {
        throw new Error(`请求失败: ${response.status}, ${response}`);
    }

    return await response.json();
}

/**
 * 请求 Descript 相关内容
 * @param {*} target 
 * @returns 
 */
export const AccessDescript = async (target) =>{

    const response = await fetch(`${apiBaseUrl}/projects/descript`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ target }),
    });

    if (!response.ok) {
        throw new Error(`请求失败: ${response.status}, ${response}`);
    }

    return await response.text();
}

/**
 * 
 * @param {*} target 
 * @returns 
 */
export const AccessData = async (target) => {
    const response = await fetch(`${apiBaseUrl}/projects/data`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ target }),
    });

    if (!response.ok) {
        throw new Error(`请求失败: ${response.status}, ${response}`);
    }

    return await response.blob();
}