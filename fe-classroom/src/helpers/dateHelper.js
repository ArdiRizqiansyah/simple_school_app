function formatDate(dateString) {
  // Buat objek Date dari string tanggal
  const date = new Date(dateString);
  
  // Dapatkan elemen-elemen tanggal
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0'); // Bulan di JavaScript 0-indexed
  const day = String(date.getDate()).padStart(2, '0');
  
  // Buat format baru, contoh: "YYYY-MM-DD"
  const formattedDate = `${year}-${month}-${day}`;
  
  return formattedDate;
}

function formatDateToString(date) {
    return new Date(date).toLocaleDateString('id-ID', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    });
}

export {
    formatDate,
    formatDateToString
}