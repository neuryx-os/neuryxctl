// Package gpu detectará as GPUs presentes e fará o setup do driver.
//
// Conforme o ADR 0004 (research/adr/0004-driver-nvidia-no-primeiro-boot.md),
// o driver NVIDIA e o CUDA não fazem parte da imagem: são instalados no
// primeiro boot a partir do repositório oficial da NVIDIA para openSUSE.
//
// Comandos previstos: "gpu detect" e "gpu setup". Ainda não implementado.
package gpu
