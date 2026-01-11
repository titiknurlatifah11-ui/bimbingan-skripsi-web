def trigger_buffer_error():
    # BufferError terjadi saat operasi buffer gagal.
    try:
        raise BufferError("Contoh buffer error")
    except BufferError as e:
        print(f"Caught a BufferError: {e}")
