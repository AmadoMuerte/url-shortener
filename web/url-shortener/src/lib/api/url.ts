import {PopUpData} from "../../types.ts";


export  const addUrl = async (login:  string, password: string, urlData: PopUpData) => {
    const response = await fetch("http://localhost:8082/url", {
        method: 'post',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Basic ' + btoa(`${login}:${password}`)
        },
        body: JSON.stringify({
            alias: urlData.alias,
            url: urlData.url
        }),
    })
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return await response.json();
}

export const editUrl = async (login:  string, password: string, urlData: PopUpData, id: number) => {
    const response = await fetch(`http://localhost:8082/url/${id}`, {
        method: 'put',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Basic ' + btoa(`${login}:${password}`)
        },
        body: JSON.stringify({
            alias: urlData.alias,
            url: urlData.url
        }),
    })
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return await response.json();
}

export const deleteUrl = async (login:  string, password: string, id: number) => {
    const response = await fetch(`http://localhost:8082/url/${id}`, {
        method: 'delete',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Basic ' + btoa(`${login}:${password}`)
        },
    })
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return await response.json();
}