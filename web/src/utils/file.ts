export const loadDataUrlAsync = (f: File) => {
  return new Promise<string>((resolve, reject) => {
    const fileReader = new FileReader()
    fileReader.onload = (ev) => {
      if (ev.target && ev.target.result) {
        const r = ev.target.result
        if (typeof r == 'string') {
          resolve(r)
        } else {
          reject()
        }
      } else {
        reject()
      }
    }
    fileReader.onerror = () => {
      reject()
    }
    fileReader.readAsDataURL(f)
  })
}
