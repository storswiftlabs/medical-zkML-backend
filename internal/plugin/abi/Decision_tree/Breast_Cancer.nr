// Code generated from Python2Noir
// inputs quantization scale reciprocal: 255
// inputs quantization zero-point: 0
// quantize_type: uint8
fn main(inputs : [i32;9],) -> pub u3 {
	if inputs[3] <= 16 {
		if inputs[2] <= 89 {
			if inputs[1] <= 191 {
				0
			} else {
				if inputs[2] <= 64 {
					if inputs[7] <= 223 {
						0
					} else {
						if inputs[0] <= 77 {
							0
						} else {
							0
						}
					}
				} else {
					if inputs[7] <= 32 {
						1
					} else {
						if inputs[7] <= 128 {
							if inputs[0] <= 128 {
								if inputs[6] <= 128 {
									0
								} else {
									1
								}
							} else {
								if inputs[6] <= 128 {
									1
								} else {
									0
								}
							}
						} else {
							0
						}
					}
				}
			}
		} else {
			if inputs[8] <= 128 {
				if inputs[7] <= 32 {
					if inputs[0] <= 77 {
						1
					} else {
						if inputs[0] <= 128 {
							0
						} else {
							if inputs[2] <= 204 {
								1
							} else {
								0
							}
						}
					}
				} else {
					if inputs[2] <= 217 {
						if inputs[2] <= 166 {
							if inputs[7] <= 159 {
								if inputs[4] <= 128 {
									1
								} else {
									if inputs[0] <= 179 {
										if inputs[6] <= 128 {
											1
										} else {
											if inputs[0] <= 128 {
												0
											} else {
												1
											}
										}
									} else {
										0
									}
								}
							} else {
								1
							}
						} else {
							0
						}
					} else {
						1
					}
				}
			} else {
				if inputs[0] <= 77 {
					if inputs[2] <= 140 {
						0
					} else {
						if inputs[7] <= 96 {
							if inputs[2] <= 166 {
								if inputs[7] <= 32 {
									1
								} else {
									0
								}
							} else {
								1
							}
						} else {
							0
						}
					}
				} else {
					if inputs[7] <= 96 {
						if inputs[2] <= 230 {
							if inputs[0] <= 128 {
								if inputs[2] <= 166 {
									if inputs[2] <= 140 {
										if inputs[2] <= 115 {
											if inputs[1] <= 191 {
												if inputs[7] <= 32 {
													1
												} else {
													0
												}
											} else {
												if inputs[6] <= 128 {
													if inputs[7] <= 32 {
														0
													} else {
														0
													}
												} else {
													0
												}
											}
										} else {
											if inputs[1] <= 191 {
												0
											} else {
												if inputs[6] <= 128 {
													1
												} else {
													if inputs[7] <= 32 {
														0
													} else {
														0
													}
												}
											}
										}
									} else {
										0
									}
								} else {
									if inputs[7] <= 32 {
										if inputs[2] <= 191 {
											1
										} else {
											0
										}
									} else {
										if inputs[2] <= 191 {
											0
										} else {
											if inputs[6] <= 128 {
												1
											} else {
												0
											}
										}
									}
								}
							} else {
								if inputs[2] <= 166 {
									if inputs[7] <= 32 {
										if inputs[6] <= 128 {
											if inputs[2] <= 140 {
												if inputs[4] <= 128 {
													0
												} else {
													if inputs[2] <= 115 {
														if inputs[0] <= 179 {
															0
														} else {
															0
														}
													} else {
														1
													}
												}
											} else {
												0
											}
										} else {
											0
										}
									} else {
										0
									}
								} else {
									if inputs[0] <= 179 {
										if inputs[7] <= 32 {
											0
										} else {
											if inputs[2] <= 191 {
												0
											} else {
												0
											}
										}
									} else {
										1
									}
								}
							}
						} else {
							1
						}
					} else {
						if inputs[7] <= 159 {
							if inputs[1] <= 191 {
								0
							} else {
								if inputs[2] <= 115 {
									0
								} else {
									if inputs[2] <= 166 {
										1
									} else {
										0
									}
								}
							}
						} else {
							if inputs[1] <= 191 {
								if inputs[2] <= 128 {
									1
								} else {
									0
								}
							} else {
								0
							}
						}
					}
				}
			}
		}
	} else {
		if inputs[6] <= 128 {
			if inputs[0] <= 128 {
				1
			} else {
				if inputs[8] <= 128 {
					if inputs[3] <= 48 {
						if inputs[2] <= 140 {
							0
						} else {
							0
						}
					} else {
						1
					}
				} else {
					if inputs[2] <= 191 {
						if inputs[7] <= 128 {
							if inputs[3] <= 112 {
								if inputs[2] <= 166 {
									if inputs[4] <= 128 {
										0
									} else {
										1
									}
								} else {
									1
								}
							} else {
								0
							}
						} else {
							1
						}
					} else {
						0
					}
				}
			}
		} else {
			if inputs[2] <= 140 {
				if inputs[0] <= 179 {
					if inputs[4] <= 128 {
						if inputs[7] <= 64 {
							if inputs[3] <= 80 {
								0
							} else {
								if inputs[3] <= 128 {
									1
								} else {
									0
								}
							}
						} else {
							1
						}
					} else {
						0
					}
				} else {
					1
				}
			} else {
				if inputs[4] <= 128 {
					if inputs[8] <= 128 {
						0
					} else {
						if inputs[2] <= 166 {
							if inputs[7] <= 32 {
								1
							} else {
								0
							}
						} else {
							1
						}
					}
				} else {
					1
				}
			}
		}
	}
}