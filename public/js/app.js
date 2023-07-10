// scrape form 
document.addEventListener('DOMContentLoaded', (event) => {
    document.getElementById('myForm').addEventListener('submit', async function (e) {
        e.preventDefault();

        const urlInput = document.getElementById('url');
        const fileTypes = document.getElementsByName('generateFile');
        const urlError = document.getElementById('urlError');
        const fileTypeError = document.getElementById('fileTypeError');
        const downloadLink = document.getElementById('downloadLink');

        urlError.textContent = '';
        fileTypeError.textContent = '';
        downloadLink.style.display = 'none';

        let selectedFileType = null;

        for (const fileType of fileTypes) {
            if (fileType.checked) {
                selectedFileType = fileType.value;
                break;
            }
        }

        if (!urlInput.value) {
            urlError.textContent = 'URL is required';
            return;
        }

        if (!selectedFileType) {
            fileTypeError.textContent = 'Please select a file type';
            return;
        }

        const data = {
            url: urlInput.value,
            generateFile: selectedFileType
        };

        try {
            const response = await fetch('http://localhost:8080/api/v1/process', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });

            if (!response.ok) {
                throw new Error('Error: ' + response.statusText);
            }

            const { filePath } = await response.json();
            const encodedFilePath = encodeURIComponent(filePath);
            const downloadLinkURL = `http://localhost:8080/api/v1/download?filePath=${encodedFilePath}`;
            downloadLink.href = downloadLinkURL;
            downloadLink.style.display = 'block';

            downloadLink.addEventListener('click', function() {
                // After clicking the link, hide it immediately
                downloadLink.style.display = 'none';
            }, { once: true }); // Listener is removed after one-time use

            
        } catch (error) {
            alert('Error: ' + error.message);
        }
    });
});