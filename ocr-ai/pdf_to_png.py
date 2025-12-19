import fitz  # PyMuPDF
from pathlib import Path
from typing import List


def pdf_to_png(
    pdf_path: str,
    output_dir: str,
    zoom: float = 3.0
) -> List[str]:
    """
    将 PDF 每一页转为 PNG 图片

    :param pdf_path: PDF 文件路径
    :param output_dir: PNG 输出目录
    :param zoom: 缩放比例，3.0 ≈ 216 DPI（推荐 OCR 使用）
    :return: PNG 文件路径列表
    """
    pdf_path = Path(pdf_path)
    output_dir = Path(output_dir)
    output_dir.mkdir(parents=True, exist_ok=True)

    doc = fitz.open(pdf_path)
    mat = fitz.Matrix(zoom, zoom)

    png_paths = []

    for i, page in enumerate(doc):
        pix = page.get_pixmap(matrix=mat, alpha=False)
        png_path = output_dir / f"{pdf_path.stem}_page_{i + 1}.png"
        pix.save(str(png_path))
        png_paths.append(str(png_path))

    doc.close()
    return png_paths
