from PIL import Image
import fitz  # PyMuPDF
from pathlib import Path
from typing import List


def pdf_to_png(
    pdf_path: str,
    output_dir: str,
    zoom: float = 3.0
) -> List[str]:
    """
    PDF 转 PNG（PDF 预处理）
    """
    output_dir = Path(output_dir)
    output_dir.mkdir(parents=True, exist_ok=True)

    png_paths = []

    # 打开PDF文档
    pdf_document = fitz.open(pdf_path)

    for page_num in range(pdf_document.page_count):
        # 获取页面
        page = pdf_document[page_num]

        # 设置缩放比例以获得更高分辨率的图像
        mat = fitz.Matrix(zoom, zoom)

        # 渲染页面为图像
        pix = page.get_pixmap(matrix=mat)

        # 保存为PNG文件
        png_path = output_dir / f"page_{page_num + 1}.png"
        pix.save(str(png_path))

        png_paths.append(str(png_path))

    # 关闭PDF文档
    pdf_document.close()

    return png_paths


def convert_image_to_jpg(
    image_path: str,
    output_dir: str,
    quality: int = 95
) -> str:
    """
    将单个图片文件转换为JPG格式
    支持多种输入格式（JPG, PNG, WebP, TIFF等）
    """
    output_dir = Path(output_dir)
    output_dir.mkdir(parents=True, exist_ok=True)

    img = Image.open(image_path)

    # 转换为RGB模式（如果需要）
    if img.mode in ("RGBA", "LA", "P"):
        img = img.convert("RGB")

    # 生成输出文件路径
    output_path = output_dir / f"{Path(image_path).stem}.jpg"
    img.save(str(output_path), "JPEG", quality=quality)

    return str(output_path)


def png_to_jpg(
    png_paths: List[str],
    output_dir: str,
    quality: int = 95
) -> List[str]:
    """
    PNG 转 JPG（统一 OCR 输入格式）
    """
    output_dir = Path(output_dir)
    output_dir.mkdir(parents=True, exist_ok=True)

    jpg_paths = []

    for png_path in png_paths:
        png_path = Path(png_path)
        img = Image.open(png_path).convert("RGB")

        jpg_path = output_dir / f"{png_path.stem}.jpg"
        img.save(jpg_path, "JPEG", quality=quality)

        jpg_paths.append(str(jpg_path))

    return jpg_paths
