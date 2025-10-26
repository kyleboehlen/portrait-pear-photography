const buildUrl = (photoId, variant) => {
    const accountHash = import.meta.env.VITE_CLOUDFLARE_ACCOUNT_HASH
    return `https://imagedelivery.net/${accountHash}/${photoId}/${variant}`
}

const fullResUrl = (photoId) => {
    return buildUrl(photoId, 'fullres')
}

const gridUrl = (photoId) => {
    return buildUrl(photoId, 'grid')
}

const previewUrl = (photoId) => {
    return buildUrl(photoId, 'preview')
}

export const usePhotoUtils = () => {
    return {fullResUrl, gridUrl, previewUrl}
}