import React from "react";

export type URLInfo = {
    id: number;
    alias: string;
    url: string;
    address: string;
}

export type PopUpConfig = {
    popupIsOpen: boolean
    setPopupIsOpen: React.Dispatch<React.SetStateAction<boolean>>
    setUpdating: React.Dispatch<React.SetStateAction<boolean>>
    setPopUpData: React.Dispatch<React.SetStateAction<PopUpData>>
    popUpData: PopUpData
    editId: number
    setEditId: React.Dispatch<React.SetStateAction<number>>
}

export type PopUpData = {
    url: string
    alias: string
}

export interface UrlResponse {
    status: string
    data: URLInfo[]
    address: string
}